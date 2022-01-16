package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const GameLibraryTable = "game_library"

// tgcon - used to generate constants for each field's tag
type GameLibrary struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}

func CreateGameLibrary(t Transaction, model GameLibrary) error {
	_, err := t.InsertInto(GameLibraryTable).
		Columns(GameLibraryIDGameDb, GameLibraryIDClientDb).
		Record(model).
		Exec()

	return err
}

func ReadGameLibraryByClientID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameLibraryTable).
		Join(GameTable, GameLibraryTable+"."+GameLibraryIDGameDb+" = "+GameTable+"."+GameIDDb).
		Where(GameLibraryTable+"."+GameLibraryIDClientDb+" = ?", id).
		Load(&objects)

	return
}

func ReadGameLibraryByID(t Transaction, gameID, clientID uuid.UUID) (object GameLibrary, ok bool, err error) {
	err = t.Select("*").
		From(GameLibraryTable).
		Where(GameLibraryIDGameDb+" = ?", gameID).
		Where(GameLibraryIDClientDb+" = ?", clientID).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
