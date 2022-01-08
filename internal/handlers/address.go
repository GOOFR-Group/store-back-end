package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerAddress = "address"

// GetAddress handles the /address Get endpoint
func (*StoreImpl) GetAddress(w http.ResponseWriter, r *http.Request, params oapi.GetAddressParams) {
	response, err := core.GetAddress(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerAddress, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerAddress, err)
		return
	default:
		writeInternalServerError(w, handlerAddress, err)
		return
	}

	writeOK(w, handlerAddress, response)
}

// PutAddress handles the /address Put endpoint
func (*StoreImpl) PutAddress(w http.ResponseWriter, r *http.Request, params oapi.PutAddressParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerAddress, err)
		return
	}

	req := oapi.PutAddressJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerAddress, err)
		return
	}

	response, err := core.PutAddress(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerAddress, fmt.Sprintf(ErrAddressNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerAddress, err)
		return
	}

	writeOK(w, handlerAddress, response)
}
