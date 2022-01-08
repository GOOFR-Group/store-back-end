package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// GetInvoice gets the client's invoice history
func GetInvoice(params oapi.GetInvoiceParams) ([]oapi.InvoiceSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []invoice

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		var invoiceHeaders []storage.InvoiceHeader

		if invoiceHeaders, err = storage.ReadInvoiceHeadersByClientID(tx, client.ID); err != nil {
			return err
		}

		var invoiceGames []storage.InvoiceGame

		for _, header := range invoiceHeaders {
			if invoiceGames, err = storage.ReadInvoiceGamesByInvoiceID(tx, header.IDInvoice); err != nil {
				return err
			}

			objects = append(objects, invoice{
				header: header,
				games:  invoiceGames,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getInvoicesFromModel(objects), nil
}

type invoice struct {
	header storage.InvoiceHeader
	games  []storage.InvoiceGame
}

func getInvoiceFromModel(model invoice) oapi.InvoiceSchema {
	var games []oapi.InvoiceGameSchema
	for _, g := range model.games {
		games = append(games, oapi.InvoiceGameSchema{
			IdGame:   g.IDGame.String(),
			Price:    g.Price,
			Discount: g.Discount,
		})
	}

	return oapi.InvoiceSchema{
		Id:           model.header.IDInvoice.String(),
		IdClient:     model.header.IDClient.String(),
		PurchaseDate: model.header.PurchaseDate,
		VatId:        model.header.VatID,
		Games:        games,
	}
}

func getInvoicesFromModel(model []invoice) []oapi.InvoiceSchema {
	array := make([]oapi.InvoiceSchema, len(model))
	for i, m := range model {
		array[i] = getInvoiceFromModel(m)
	}
	return array
}
