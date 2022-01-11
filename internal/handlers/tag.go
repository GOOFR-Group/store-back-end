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
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrParsingRequest, err.Error())}
		if err = writeResponse(w, response, http.StatusBadRequest, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	}

	req := oapi.PostTagJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrParsingRequest, err.Error())}
		if err = writeResponse(w, response, http.StatusBadRequest, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	}

	err = core.PostTag(req)
	switch err {
	case nil:
	default:
		logInternalError(handlerTag, ErrInternalServer, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetTag handles the /tag Get endpoint
func (*StoreImpl) GetTag(w http.ResponseWriter, r *http.Request, params oapi.GetTagParams) {
	response, err := core.GetTag(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrTagNotFound, *params.Id)}
		if err = writeResponse(w, response, http.StatusNotFound, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	default:
		logInternalError(handlerTag, ErrInternalServer, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = writeResponse(w, response, http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerTag, ErrWritingResponse, err)
	}
}

// PutTag handles the /tag Put endpoint
func (*StoreImpl) PutTag(w http.ResponseWriter, r *http.Request, params oapi.PutTagParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrParsingRequest, err.Error())}
		if err = writeResponse(w, response, http.StatusBadRequest, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	}

	req := oapi.PutTagJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrParsingRequest, err.Error())}
		if err = writeResponse(w, response, http.StatusBadRequest, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	}

	err = core.PutTag(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrTagNotFound, params.Id)}
		if err = writeResponse(w, response, http.StatusNotFound, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	default:
		logInternalError(handlerTag, ErrInternalServer, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTag handles the /tag Delete endpoint
func (*StoreImpl) DeleteTag(w http.ResponseWriter, r *http.Request, params oapi.DeleteTagParams) {
	response, err := core.DeleteTag(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrTagNotFound, params.Id)}
		if err = writeResponse(w, response, http.StatusNotFound, header{contentTypeHeader, contentTypeJSON}); err != nil {
			logInternalError(handlerTag, ErrWritingResponse, err)
		}
		return
	default:
		logInternalError(handlerTag, ErrInternalServer, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = writeResponse(w, response, http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerTag, ErrWritingResponse, err)
	}
}
