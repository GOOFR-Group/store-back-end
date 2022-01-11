package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerInvoice = "invoice"

// GetInvoice handles the /invoice Get endpoint
func (*StoreImpl) GetInvoice(w http.ResponseWriter, r *http.Request, params oapi.GetInvoiceParams) {
}
