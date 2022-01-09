package core

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

const searchLimit = 30

// GetSearchGame search games
func GetSearchGame(params oapi.GetSearchGameParams) ([]oapi.GameSchema, error) {
	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if objects, err = storage.ReadGamesByNameLike(tx, params.Search, searchLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

// GetSearchTag search tags
func GetSearchTag(params oapi.GetSearchTagParams) ([]oapi.TagSchema, error) {
	var objects []storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if objects, err = storage.ReadTagsByNameLike(tx, params.Search, searchLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getTagsFromModel(objects), nil
}

// PostSearchHistory adds a search to the client's history
func PostSearchHistory(params oapi.PostSearchHistoryParams) error {
	var idClient uuid.UUID
	var idGame uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.ClientID); err != nil {
		return err
	}

	if idGame, err = uuid.Parse(params.GameID); err != nil {
		return err
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var game storage.Game
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if game, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		var id uuid.UUID

		if id, err = uuid.NewRandom(); err != nil {
			return fmt.Errorf(ErrGeneratingUUID, err.Error())
		}

		if err = storage.CreateClientSearchHistory(tx, storage.ClientSearchHistory{
			ID:       id,
			IDGame:   game.ID,
			IDClient: client.ID,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetSearchHistory gets the client's search history
func GetSearchHistory(params oapi.GetSearchHistoryParams) ([]oapi.SearchSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.ClientSearchHistory

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool
		var err error

		if _, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if objects, err = storage.ReadClientSearchHistoryByClientID(tx, idClient); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getSearchsFromModel(objects), nil
}

func getSearchFromModel(model storage.ClientSearchHistory) oapi.SearchSchema {
	return oapi.SearchSchema{
		Id:       model.ID.String(),
		IdGame:   model.IDGame.String(),
		IdClient: model.IDClient.String(),
		DateTime: model.DateTime,
	}
}

func getSearchsFromModel(model []storage.ClientSearchHistory) []oapi.SearchSchema {
	array := make([]oapi.SearchSchema, len(model))
	for i, m := range model {
		array[i] = getSearchFromModel(m)
	}
	return array
}
