package storage

import (
	"time"

	"github.com/google/uuid"
)

const InvoiceGameTable = "invoice_game"

// tgcon - used to generate constants for each field's tag
type InvoiceGame struct {
	IDInvoice uuid.UUID `db:"id_invoice"`
	IDGame    uuid.UUID `db:"id_game"`
	Price     time.Time `db:"price"`
	Discount  int64     `db:"discount"`
}
