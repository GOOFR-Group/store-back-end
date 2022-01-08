package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerStore = "store"

// GetYourStore handles the /yourStore Get endpoint
func (*StoreImpl) GetYourStore(w http.ResponseWriter, r *http.Request, params oapi.GetYourStoreParams) {
}

// GetNewStore handles the /newStore Get endpoint
func (*StoreImpl) GetNewStore(w http.ResponseWriter, r *http.Request, params oapi.GetNewStoreParams) {
}

// GetNoteworthyStore handles the /noteworthyStore Get endpoint
func (*StoreImpl) GetNoteworthyStore(w http.ResponseWriter, r *http.Request, params oapi.GetNoteworthyStoreParams) {
}
