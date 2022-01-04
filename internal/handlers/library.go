package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerLibrary = "library"

// GetLibrary handles the /library Get endpoint
func (*StoreImpl) GetLibrary(w http.ResponseWriter, r *http.Request, params oapi.GetLibraryParams) {
}
