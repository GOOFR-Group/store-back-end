package storage

import (
	"time"

	"github.com/gocraft/dbr/v2"
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

func CreateInvoiceHeader(t Transaction, model InvoiceHeader) error {
	_, err := t.InsertInto(InvoiceHeaderTable).
		Columns(InvoiceHeaderIDInvoiceDb, InvoiceHeaderIDClientDb, InvoiceHeaderPurchaseDateDb, InvoiceHeaderVatIDDb).
		Record(model).
		Exec()

	return err
}

func ReadInvoiceHeadersByClientID(t Transaction, id uuid.UUID) (objects []InvoiceHeader, err error) {
	_, err = t.Select(InvoiceHeaderTable).
		From(InvoiceHeaderTable).
		Where(InvoiceHeaderIDClientDb+" = ?", id).
		OrderDesc(InvoiceHeaderPurchaseDateDb).
		Load(&objects)

	return
}
func ReadInvoiceHeaderByID(t Transaction, id uuid.UUID) (object InvoiceHeader, ok bool, err error) {
	err = t.Select("*").
		From(InvoiceHeaderTable).
		Where(InvoiceHeaderIDInvoiceDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
