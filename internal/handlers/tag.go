package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/core"
	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerTag = "tag"

// PostTag handles the /tag Post endpoint
func (*StoreImpl) PostTag(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerTag, err)
		return
	}

	req := oapi.PostTagJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerTag, err)
		return
	}

	err = core.PostTag(req)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerTag, err)
		return
	}

	writeCreated(w)
}

// GetTag handles the /tag Get endpoint
func (*StoreImpl) GetTag(w http.ResponseWriter, r *http.Request, params oapi.GetTagParams) {
	response, err := core.GetTag(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerTag, fmt.Sprintf(ErrTagNotFound, *params.Id))
		return
	default:
		writeInternalServerError(w, handlerTag, err)
		return
	}

	writeOK(w, handlerTag, response)
}

// PutTag handles the /tag Put endpoint
func (*StoreImpl) PutTag(w http.ResponseWriter, r *http.Request, params oapi.PutTagParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerTag, err)
		return
	}

	req := oapi.PutTagJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerTag, err)
		return
	}

	err = core.PutTag(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerTag, fmt.Sprintf(ErrTagNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerTag, err)
		return
	}

	writeNoContent(w)
}

// DeleteTag handles the /tag Delete endpoint
func (*StoreImpl) DeleteTag(w http.ResponseWriter, r *http.Request, params oapi.DeleteTagParams) {
	response, err := core.DeleteTag(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerTag, fmt.Sprintf(ErrTagNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerTag, err)
		return
	}

	writeOK(w, handlerTag, response)
}
