package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const ImageGameTable = "image_game"

// tgcon - used to generate constants for each field's tag
type ImageGame struct {
	ID     uuid.UUID `db:"id"`
	IDGame uuid.UUID `db:"id_game"`
	Image  string    `db:"image"`
}

func CreateImageGame(t Transaction, model ImageGame) error {
	_, err := t.InsertInto(ImageGameTable).
		Columns(ImageGameIDDb, ImageGameIDGameDb, ImageGameImageDb).
		Record(model).
		Exec()

	return err
}

func ReadImageGamesByGameID(t Transaction, id uuid.UUID) (objects []ImageGame, err error) {
	_, err = t.Select("*").
		From(ImageGameTable).
		Where(ImageGameIDGameDb+" = ?", id).
		Load(&objects)

	return
}

func ReadImageGameByID(t Transaction, id uuid.UUID) (object ImageGame, ok bool, err error) {
	err = t.Select("*").
		From(ImageGameTable).
		Where(ImageGameIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func DeleteImageGameByID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(ImageGameTable).
		Where(ImageGameIDDb+" = ?", id).
		Exec()

	return err
}
