package storage

import (
	"time"

	"github.com/google/uuid"
)

const InvoiceHeaderTable = "invoice_header"

// tgcon - used to generate constants for each field's tag
type InvoiceHeader struct {
	IDInvoice    uuid.UUID `db:"id_invoice"`
	IDClient     uuid.UUID `db:"id_client"`
	PurchaseDate time.Time `db:"purchase_date"`
	VatID        int64     `db:"vat_id"`
}
