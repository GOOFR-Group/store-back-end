package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const WishlistTable = "wishlist"

// tgcon - used to generate constants for each field's tag
type Wishlist struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}

func CreateWishlist(t Transaction, model Wishlist) error {
	_, err := t.InsertInto(WishlistTable).
		Columns(WishlistIDGameDb, WishlistIDClientDb).
		Record(model).
		Exec()

	return err
}

func ReadWishlistGamesByClientID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select("*").
		From(WishlistTable).
		Join(GameTable, WishlistTable+"."+WishlistIDGameDb+" = "+GameTable+"."+GameIDDb).
		Where(WishlistTable+"."+WishlistIDClientDb+" = ?", id).
		Load(&objects)

	return
}

func ReadWishlistByID(t Transaction, gameID, clientID uuid.UUID) (object Wishlist, ok bool, err error) {
	err = t.Select("*").
		From(WishlistTable).
		Where(WishlistIDGameDb+" = ?", gameID).
		Where(WishlistIDClientDb+" = ?", clientID).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadClientEmailsByGameInWishlis(t Transaction, gameID uuid.UUID) (objects []string, err error) {
	_, err = t.Select(AccessTable+"."+AccessEmailDb).
		From(AccessTable).
		Join(ClientTable, AccessTable+"."+AccessIDClientDb+" = "+ClientTable+"."+ClientIDDb).
		Join(WishlistTable, ClientTable+"."+ClientIDDb+" = "+WishlistTable+"."+WishlistIDClientDb).
		Where(WishlistTable+"."+WishlistIDGameDb+" = ?", gameID).
		Load(&objects)

	return
}

func DeleteWishlistByID(t Transaction, gameID, clientID uuid.UUID) error {
	_, err := t.DeleteFrom(WishlistTable).
		Where(WishlistIDGameDb+" = ?", gameID).
		Where(WishlistIDClientDb+" = ?", clientID).
		Exec()

	return err
}

func DeleteWishlistByClientID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(WishlistTable).
		Where(WishlistIDClientDb+" = ?", id).
		Exec()

	return err
}
