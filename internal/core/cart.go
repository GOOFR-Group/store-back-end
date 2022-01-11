package core

import (
	"fmt"
	"net/smtp"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/conf"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/goofr-group/store-back-end/internal/utils/mathf"
	"github.com/google/uuid"
)

const timeLayout = "02/01/2006"

// PostCart adds a game to the client's cart
func PostCart(params oapi.PostCartParams) error {
	var idClient uuid.UUID
	var idGame uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.ClientID); err != nil {
		return err
	}

	if idGame, err = uuid.Parse(params.GameID); err != nil {
		return err
	}

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if _, ok, err = storage.ReadGameLibraryByID(tx, idGame, idClient); err != nil {
			return err
		}
		if ok {
			return ErrGameAlreadyBought
		}

		if _, ok, err = storage.ReadCartByID(tx, idGame, idClient); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		if err = storage.CreateCart(tx, storage.Cart{
			IDGame:   idGame,
			IDClient: idClient,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetCart gets all the games the client has in his cart
func GetCart(params oapi.GetCartParams) ([]oapi.GameSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if objects, err = storage.ReadCartGamesByClientID(tx, idClient); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

// DeleteCart removes a game from the client's cart
func DeleteCart(params oapi.DeleteCartParams) ([]oapi.GameSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.ClientID); err != nil {
		return nil, err
	}

	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if params.GameID == nil {
			if objects, err = storage.ReadCartGamesByClientID(tx, idClient); err != nil {
				return err
			}

			if err = storage.DeleteCartByClientID(tx, idClient); err != nil {
				return err
			}
		} else {
			var idGame uuid.UUID

			if idGame, err = uuid.Parse(*params.GameID); err != nil {
				return err
			}

			var object storage.Game

			if object, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
				return err
			}
			if !ok {
				return ErrGameNotFound
			}

			if _, ok, err = storage.ReadCartByID(tx, idGame, idClient); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			if err = storage.DeleteCartByID(tx, idGame, idClient); err != nil {
				return err
			}

			objects = append(objects, object)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

// GetCartPurchase purchases all games the client has in his cart
func GetCartPurchase(params oapi.GetCartPurchaseParams) (oapi.InvoiceSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return oapi.InvoiceSchema{}, err
	}

	var object invoice
	var objectGames []storage.Game
	var total float64
	var clientAccess storage.Access
	var clientWallet storage.Wallet

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if clientAccess, ok, err = storage.ReadAccessByClientID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		var games []storage.Game

		if games, err = storage.ReadCartGamesByClientID(tx, client.ID); err != nil {
			return err
		}
		if games == nil {
			return ErrObjectNotFound
		}

		for _, g := range games {
			total += g.Price * (1 + mathf.Clamp(g.Discount, 0, 1))
		}

		if clientWallet, ok, err = storage.ReadWalletByClientID(tx, client.ID); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if clientWallet.Balance < total {
			return ErrInvalidAmount
		}

		if err = storage.UpdateWalletByID(tx, storage.Wallet{
			ID:       clientWallet.ID,
			IDClient: clientWallet.IDClient,
			Balance:  clientWallet.Balance - total,
			Coin:     clientWallet.Coin,
		}); err != nil {
			return err
		}

		var idInvoice uuid.UUID

		if idInvoice, err = uuid.NewRandom(); err != nil {
			return fmt.Errorf(ErrGeneratingUUID, err.Error())
		}

		invoiceHeader := storage.InvoiceHeader{
			IDInvoice: idInvoice,
			IDClient:  client.ID,
			VatID:     client.VatID,
		}

		if err = storage.CreateInvoiceHeader(tx, invoiceHeader); err != nil {
			return err
		}

		var invoiceGames []storage.InvoiceGame

		for _, g := range games {
			if _, ok, err = storage.ReadGameLibraryByID(tx, g.ID, idClient); err != nil {
				return err
			}
			if ok {
				continue
			}

			if err = storage.DeleteWishlistByID(tx, g.ID, idClient); err != nil {
				return err
			}

			if err = storage.CreateGameLibrary(tx, storage.GameLibrary{
				IDGame:   g.ID,
				IDClient: idClient,
			}); err != nil {
				return err
			}

			invoiceGame := storage.InvoiceGame{
				IDInvoice: idInvoice,
				IDGame:    g.ID,
				Price:     g.Price,
				Discount:  g.Discount,
			}

			if err = storage.CreateInvoiceGame(tx, invoiceGame); err != nil {
				return err
			}

			invoiceGames = append(invoiceGames, invoiceGame)
			objectGames = append(objectGames, g)
		}

		if invoiceHeader, ok, err = storage.ReadInvoiceHeaderByID(tx, idInvoice); err != nil {
			return err
		}
		if !ok {
			return ErrInvoiceHeaderNotFound
		}

		object = invoice{
			header: invoiceHeader,
			games:  invoiceGames,
		}

		if err = storage.DeleteCartByClientID(tx, idClient); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.InvoiceSchema{}, err
	}

	title := fmt.Sprintf("GOOFR Store - Invoice - %s", object.header.PurchaseDate.Format(timeLayout))

	body := fmt.Sprintf("ID: %s\n", object.header.IDInvoice.String())
	body += fmt.Sprintf("Purchase Date: %s\n", object.header.PurchaseDate.Format(timeLayout))
	body += fmt.Sprintf("Vat ID: %d\n", object.header.VatID)
	body += "Games: \n"
	for _, g := range objectGames {
		body += fmt.Sprintf("\nName: %s\n", g.Name)
		body += fmt.Sprintf("\tState: %s\n", g.State)
		body += fmt.Sprintf("\tRelease Date: %s\n", g.ReleaseDate.Format(timeLayout))
		body += fmt.Sprintf("\tPrice: %s%.2f\n", clientWallet.Coin, g.Price)
		body += fmt.Sprintf("\tDiscount: %.2f%%\n", g.Discount*100)
	}
	body += fmt.Sprintf("\nTotal: %s%.2f", clientWallet.Coin, total)

	message := []byte(fmt.Sprintf(smtpSubject, title) + body)
	if err = smtp.SendMail(conf.SMTPAddress(), conf.SMTPAuthentication(), conf.SMTPEmailAddress(), []string{clientAccess.Email}, message); err != nil {
		return oapi.InvoiceSchema{}, err
	}

	return getInvoiceFromModel(object), nil
}
