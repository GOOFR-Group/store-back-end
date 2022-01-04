package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerAddress = "address"

// GetAddress handles the /address Get endpoint
func (*StoreImpl) GetAddress(w http.ResponseWriter, r *http.Request, params oapi.GetAddressParams) {
}

// PutAddress handles the /address Put endpoint
func (*StoreImpl) PutAddress(w http.ResponseWriter, r *http.Request, params oapi.PutAddressParams) {
}
