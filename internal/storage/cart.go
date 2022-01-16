package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const CartTable = "cart"

// tgcon - used to generate constants for each field's tag
type Cart struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}

func CreateCart(t Transaction, model Cart) error {
	_, err := t.InsertInto(CartTable).
		Columns(CartIDGameDb, CartIDClientDb).
		Record(model).
		Exec()

	return err
}

func ReadCartGamesByClientID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select("*").
		From(CartTable).
		Join(GameTable, CartTable+"."+CartIDGameDb+" = "+GameTable+"."+GameIDDb).
		Where(CartTable+"."+CartIDClientDb+" = ?", id).
		Load(&objects)

	return
}

func ReadCartByID(t Transaction, gameID, clientID uuid.UUID) (object Cart, ok bool, err error) {
	err = t.Select("*").
		From(CartTable).
		Where(CartIDGameDb+" = ?", gameID).
		Where(CartIDClientDb+" = ?", clientID).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func DeleteCartByID(t Transaction, gameID, clientID uuid.UUID) error {
	_, err := t.DeleteFrom(CartTable).
		Where(CartIDGameDb+" = ?", gameID).
		Where(CartIDClientDb+" = ?", clientID).
		Exec()

	return err
}

func DeleteCartByClientID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(CartTable).
		Where(CartIDClientDb+" = ?", id).
		Exec()

	return err
}
