package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerClient = "client"

// GetClient handles the /client Get endpoint
func (*StoreImpl) GetClient(w http.ResponseWriter, r *http.Request, params oapi.GetClientParams) {
	response, err := core.GetClient(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerClient, fmt.Sprintf(ErrClientNotFound, *params.Id))
		return
	default:
		writeInternalServerError(w, handlerClient, err)
		return
	}

	writeOK(w, handlerClient, response)
}

// PutClient handles the /client Put endpoint
func (*StoreImpl) PutClient(w http.ResponseWriter, r *http.Request, params oapi.PutClientParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerClient, err)
		return
	}

	req := oapi.PutClientJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerClient, err)
		return
	}

	response, err := core.PutClient(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerClient, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerClient, err)
		return
	}

	writeOK(w, handlerClient, response)
}

// DeleteClient handles the /client Delete endpoint
func (*StoreImpl) DeleteClient(w http.ResponseWriter, r *http.Request, params oapi.DeleteClientParams) {
	response, err := core.DeleteClient(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerClient, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerClient, err)
		return
	}

	writeOK(w, handlerClient, response)
}
