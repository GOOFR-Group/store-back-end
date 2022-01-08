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

func ReadInvoiceHeadersByClientID(t Transaction, id uuid.UUID) (objects []InvoiceHeader, err error) {
	_, err = t.Select(InvoiceHeaderTable).
		From(InvoiceHeaderTable).
		Where(InvoiceHeaderIDClientDb+" = ?", id).
		OrderDesc(InvoiceHeaderPurchaseDateDb).
		Load(&objects)

	return
}
