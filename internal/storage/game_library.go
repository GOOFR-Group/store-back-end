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
