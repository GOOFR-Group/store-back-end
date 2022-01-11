package handlers

import (
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerInvoice = "invoice"

// GetInvoice handles the /invoice Get endpoint
func (*StoreImpl) GetInvoice(w http.ResponseWriter, r *http.Request, params oapi.GetInvoiceParams) {
	response, err := core.GetInvoice(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerInvoice, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerInvoice, err)
		return
	}

	writeOK(w, handlerInvoice, response)
}
