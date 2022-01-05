package storage

import (
	"database/sql"

	"github.com/google/uuid"
)

const ReviewTable = "review"

// tgcon - used to generate constants for each field's tag
type Review struct {
	ID       uuid.UUID      `db:"id"`
	IDGame   uuid.UUID      `db:"id_game"`
	IDClient string         `db:"id_client"`
	Stars    sql.NullInt64  `db:"stars"`
	Review   sql.NullString `db:"review"`
}
