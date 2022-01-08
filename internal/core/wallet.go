package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// GetWallet gets a client's wallet
func GetWallet(params oapi.GetWalletParams) (oapi.ClientWalletSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	var object storage.Wallet

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if object, ok, err = storage.ReadWalletByClientID(tx, client.ID); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	return getWalletFromModel(object), nil
}

// PutWallet updates a wallet
func PutWallet(params oapi.PutWalletParams, req oapi.PutWalletJSONRequestBody) (oapi.ClientWalletSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	var object storage.Wallet

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if err = storage.UpdateWalletByID(tx, storage.Wallet{
			ID:      id,
			Balance: req.Balance,
			Coin:    req.Coin,
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadWalletByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	return getWalletFromModel(object), nil
}

// GetAddBalance adds balance to the client's wallet
func GetAddBalance(params oapi.GetAddBalanceParams) (oapi.ClientWalletSchema, error) {
	if params.Amount <= 0 {
		return oapi.ClientWalletSchema{}, ErrInvalidAmount
	}

	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	var object storage.Wallet

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if object, ok, err = storage.ReadWalletByClientID(tx, client.ID); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.UpdateWalletByID(tx, storage.Wallet{
			ID:      object.ID,
			Balance: object.Balance + params.Amount,
			Coin:    object.Coin,
		}); err != nil {
			return err
		}

		if object, ok, err = storage.ReadWalletByID(tx, object.ID); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientWalletSchema{}, err
	}

	return getWalletFromModel(object), nil
}

func getWalletFromModel(model storage.Wallet) oapi.ClientWalletSchema {
	id := model.ID.String()
	idClient := model.IDClient.String()
	return oapi.ClientWalletSchema{
		Id:       &id,
		IdClient: &idClient,
		Balance:  model.Balance,
		Coin:     model.Coin,
	}
}
