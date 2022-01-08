package handlers

import (
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerLibrary = "library"

// GetLibrary handles the /library Get endpoint
func (*StoreImpl) GetLibrary(w http.ResponseWriter, r *http.Request, params oapi.GetLibraryParams) {
	response, err := core.GetLibrary(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerLibrary, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerLibrary, err)
		return
	}

	writeOK(w, handlerLibrary, response)
}
