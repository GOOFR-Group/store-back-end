package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerStore = "store"

// GetYourStore handles the /yourStore Get endpoint
func (*StoreImpl) GetYourStore(w http.ResponseWriter, r *http.Request, params oapi.GetYourStoreParams) {
	response, err := core.GetYourStore(params)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerStore, err)
		return
	}

	writeOK(w, handlerStore, response)
}

// GetNewStore handles the /newStore Get endpoint
func (*StoreImpl) GetNewStore(w http.ResponseWriter, r *http.Request, params oapi.GetNewStoreParams) {
	response, err := core.GetNewStore(params)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerStore, err)
		return
	}

	writeOK(w, handlerStore, response)
}

// GetNoteworthyStore handles the /noteworthyStore Get endpoint
func (*StoreImpl) GetNoteworthyStore(w http.ResponseWriter, r *http.Request, params oapi.GetNoteworthyStoreParams) {
	response, err := core.GetNoteworthyStore(params)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerStore, err)
		return
	}

	writeOK(w, handlerStore, response)
}
