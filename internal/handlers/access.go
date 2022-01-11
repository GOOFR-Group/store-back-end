package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerAccess = "access"

// PostLogin handles the /login Post endpoint
func (*StoreImpl) PostLogin(w http.ResponseWriter, r *http.Request) {
}

// PostRegister handles the /register Post endpoint
func (*StoreImpl) PostRegister(w http.ResponseWriter, r *http.Request) {
}

// PutAccess handles the /access Put endpoint
func (*StoreImpl) PutAccess(w http.ResponseWriter, r *http.Request, params oapi.PutAccessParams) {
}
