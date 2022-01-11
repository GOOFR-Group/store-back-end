package core

import (
	"database/sql"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// GetAddress gets a client's address
func GetAddress(params oapi.GetAddressParams) (oapi.ClientAddressSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientAddressSchema{}, err
	}

	var object storage.Address

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if object, ok, err = storage.ReadAddressByClientID(tx, client.ID); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientAddressSchema{}, err
	}

	return getAddressFromModel(object), nil
}

// PutAddress updates an address
func PutAddress(params oapi.PutAddressParams, req oapi.PutAddressJSONRequestBody) (oapi.ClientAddressSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientAddressSchema{}, err
	}

	var object storage.Address

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var doorNumber string
		var doorNumberOK bool
		if req.DoorNumber != nil {
			doorNumber = *req.DoorNumber
			doorNumberOK = true
		}

		if err = storage.UpdateAddressByID(tx, storage.Address{
			ID:     id,
			Street: req.Street,
			DoorNumber: sql.NullString{
				String: doorNumber,
				Valid:  doorNumberOK,
			},
			ZipCode: req.ZipCode,
			City:    req.City,
			Country: req.Country,
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadAddressByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientAddressSchema{}, err
	}

	return getAddressFromModel(object), nil
}

func getAddressFromModel(model storage.Address) oapi.ClientAddressSchema {
	id := model.ID.String()
	idClient := model.IDClient.String()
	return oapi.ClientAddressSchema{
		Id:         &id,
		IdClient:   &idClient,
		Street:     model.Street,
		DoorNumber: &model.DoorNumber.String,
		ZipCode:    model.ZipCode,
		City:       model.City,
		Country:    model.Country,
	}
}
