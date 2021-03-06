package core

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// PostPublisher creates a new publisher
func PostPublisher(req oapi.PostPublisherJSONRequestBody) (oapi.PublisherSchema, error) {
	var id uuid.UUID
	var object storage.Publisher
	var err error

	if id, err = uuid.NewRandom(); err != nil {
		return oapi.PublisherSchema{}, fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if err = storage.CreatePublisher(tx, storage.Publisher{
			ID:          id,
			Name:        req.Name,
			CoverImage:  req.CoverImage,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadPublisherByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.PublisherSchema{}, err
	}

	return getPublisherFromModel(object), nil
}

// GetPublisher gets a publisher
func GetPublisher(params oapi.GetPublisherParams) ([]oapi.PublisherSchema, error) {
	var objects []storage.Publisher

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if params.Id == nil {
			if objects, err = storage.ReadPublishers(tx); err != nil {
				return err
			}
		} else {
			var id uuid.UUID

			if id, err = uuid.Parse(*params.Id); err != nil {
				return err
			}

			var object storage.Publisher
			var ok bool

			if object, ok, err = storage.ReadPublisherByID(tx, id); err != nil {
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

	return getPublishersFromModel(objects), nil
}

// PutPublisher updates a publisher
func PutPublisher(params oapi.PutPublisherParams, req oapi.PutPublisherJSONRequestBody) (oapi.PublisherSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.PublisherSchema{}, err
	}

	var object storage.Publisher

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if err = storage.UpdatePublisherByID(tx, storage.Publisher{
			ID:          id,
			Name:        req.Name,
			CoverImage:  req.CoverImage,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadPublisherByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.PublisherSchema{}, err
	}

	return getPublisherFromModel(object), nil
}

// DeletePublisher deletes a publisher
func DeletePublisher(params oapi.DeletePublisherParams) (oapi.PublisherSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.PublisherSchema{}, err
	}

	var object storage.Publisher

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadPublisherByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeletePublisherByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.PublisherSchema{}, err
	}

	return getPublisherFromModel(object), nil
}

// GetPublisherGames gets all the publisher's games
func GetPublisherGames(params oapi.GetPublisherGamesParams) ([]oapi.GameSchema, error) {
	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var id uuid.UUID
		var err error

		if id, err = uuid.Parse(params.Id); err != nil {
			return err
		}

		var ok bool

		if _, ok, err = storage.ReadPublisherByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if objects, err = storage.ReadGamesByPublisherID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

func getPublisherFromModel(model storage.Publisher) oapi.PublisherSchema {
	id := model.ID.String()
	return oapi.PublisherSchema{
		Id:          &id,
		Name:        model.Name,
		CoverImage:  model.CoverImage,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
	}
}

func getPublishersFromModel(model []storage.Publisher) []oapi.PublisherSchema {
	array := make([]oapi.PublisherSchema, len(model))
	for i, m := range model {
		array[i] = getPublisherFromModel(m)
	}
	return array
}
