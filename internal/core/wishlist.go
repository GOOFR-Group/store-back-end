package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// PostWishlist adds a game to the client's wishlist
func PostWishlist(params oapi.PostWishlistParams) error {
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

		if _, ok, err = storage.ReadWishlistByID(tx, idGame, idClient); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		if err = storage.CreateWishlist(tx, storage.Wishlist{
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

// GetWishlist gets all the games the client has in his wishlist
func GetWishlist(params oapi.GetWishlistParams) ([]oapi.GameSchema, error) {
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

		if objects, err = storage.ReadWishlistGamesByClientID(tx, idClient); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

// DeleteWishlist removes a game from the client's wishlist
func DeleteWishlist(params oapi.DeleteWishlistParams) ([]oapi.GameSchema, error) {
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
			if objects, err = storage.ReadWishlistGamesByClientID(tx, idClient); err != nil {
				return err
			}

			if err = storage.DeleteWishlistByClientID(tx, idClient); err != nil {
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

			if _, ok, err = storage.ReadWishlistByID(tx, idGame, idClient); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			if err = storage.DeleteWishlistByID(tx, idGame, idClient); err != nil {
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
