package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// GetLibrary gets all the games the client has in his library
func GetLibrary(params oapi.GetLibraryParams) ([]oapi.GameSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if objects, err = storage.ReadGameLibraryByClientID(tx, client.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}
