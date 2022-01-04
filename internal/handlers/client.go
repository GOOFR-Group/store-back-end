package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerClient = "client"

// GetClient handles the /client Get endpoint
func (*StoreImpl) GetClient(w http.ResponseWriter, r *http.Request, params oapi.GetClientParams) {
}

// PutClient handles the /client Put endpoint
func (*StoreImpl) PutClient(w http.ResponseWriter, r *http.Request, params oapi.PutClientParams) {
}

// DeleteClient handles the /client Delete endpoint
func (*StoreImpl) DeleteClient(w http.ResponseWriter, r *http.Request, params oapi.DeleteClientParams) {
}
