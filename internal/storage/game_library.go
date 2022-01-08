package storage

import (
	"github.com/google/uuid"
)

const GameLibraryTable = "game_library"

// tgcon - used to generate constants for each field's tag
type GameLibrary struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}

func ReadGameLibraryByClientID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameLibraryTable).
		Join(GameTable, GameLibraryTable+"."+GameLibraryIDGameDb+" = "+GameTable+"."+GameIDDb).
		Where(GameLibraryTable+"."+GameLibraryIDClientDb+" = ?", id).
		Load(&objects)

	return
}
