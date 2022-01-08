package storage

import (
	"github.com/google/uuid"
)

const InvoiceGameTable = "invoice_game"

// tgcon - used to generate constants for each field's tag
type InvoiceGame struct {
	IDInvoice uuid.UUID `db:"id_invoice"`
	IDGame    uuid.UUID `db:"id_game"`
	Price     float64   `db:"price"`
	Discount  float64   `db:"discount"`
}

func ReadInvoiceGamesByInvoiceID(t Transaction, id uuid.UUID) (objects []InvoiceGame, err error) {
	_, err = t.Select(InvoiceGameTable).
		From(InvoiceGameTable).
		Where(InvoiceGameIDInvoiceDb+" = ?", id).
		Load(&objects)

	return
}
