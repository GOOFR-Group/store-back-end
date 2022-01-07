package core

import (
	"github.com/GOOFR-Group/store-back-end/internal/oapi"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

// GetClient gets a client
func GetClient(params oapi.GetClientParams) ([]oapi.ClientSchema, error) {
	var objects []storage.Client

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if params.Id == nil {
			if objects, err = storage.ReadClients(tx); err != nil {
				return err
			}
		} else {
			var id uuid.UUID

			if id, err = uuid.Parse(*params.Id); err != nil {
				return err
			}

			var object storage.Client
			var ok bool

			if object, ok, err = storage.ReadClientByID(tx, id); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			objects = append(objects, object)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getClientsFromModel(objects), nil
}

// PutClient updates a client
func PutClient(params oapi.PutClientParams, req oapi.PutClientJSONRequestBody) (oapi.ClientSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientSchema{}, err
	}

	var object storage.Client

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if err = storage.UpdateClientByID(tx, storage.Client{
			ID:          id,
			Name:        req.Name,
			Surname:     req.Surname,
			Picture:     req.Picture,
			Birthdate:   req.Birthdate.Time,
			PhoneNumber: req.PhoneNumber,
			VatID:       req.VatId,
			Active:      req.Active,
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadClientByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientSchema{}, err
	}

	return getClientFromModel(object), nil
}

// DeleteClient deletes a client
func DeleteClient(params oapi.DeleteClientParams) (oapi.ClientSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ClientSchema{}, err
	}

	var object storage.Client

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadClientByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteClientByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.ClientSchema{}, err
	}

	return getClientFromModel(object), nil
}

func getClientFromModel(model storage.Client) oapi.ClientSchema {
	id := model.ID.String()
	return oapi.ClientSchema{
		Id:          &id,
		Name:        model.Name,
		Surname:     model.Surname,
		Picture:     model.Picture,
		Birthdate:   openapi_types.Date{Time: model.Birthdate},
		PhoneNumber: model.PhoneNumber,
		VatId:       model.VatID,
		Active:      model.Active,
	}
}

func getClientsFromModel(model []storage.Client) []oapi.ClientSchema {
	array := make([]oapi.ClientSchema, len(model))
	for i, m := range model {
		array[i] = getClientFromModel(m)
	}
	return array
}
