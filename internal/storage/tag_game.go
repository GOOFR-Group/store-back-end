package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const TagGameTable = "tag_game"

// tgcon - used to generate constants for each field's tag
type TagGame struct {
	IDTag  uuid.UUID `db:"id_tag"`
	IDGame uuid.UUID `db:"id_game"`
}

func CreateTagGame(t Transaction, tagID, gameID uuid.UUID) error {
	_, err := t.InsertInto(TagGameTable).
		Pair(TagGameIDTagDb, tagID).
		Pair(TagGameIDGameDb, gameID).
		Exec()

	return err
}

func ReadTagGameByID(t Transaction, tagID, gameID uuid.UUID) (object TagGame, ok bool, err error) {
	err = t.Select("*").
		From(TagGameTable).
		Where(TagGameIDTagDb+" = ?", tagID).
		Where(TagGameIDGameDb+" = ?", gameID).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func DeleteTagGameByID(t Transaction, tagID, gameID uuid.UUID) (ok bool, err error) {
	_, err = t.DeleteFrom(TagGameTable).
		Where(TagGameIDTagDb+" = ?", tagID).
		Where(TagGameIDGameDb+" = ?", gameID).
		Exec()

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
