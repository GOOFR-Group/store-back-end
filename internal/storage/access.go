package storage

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const AccessTable = "access"

// tgcon - used to generate constants for each field's tag
type Access struct {
	ID        uuid.UUID      `db:"id"`
	IDClient  uuid.UUID      `db:"id_client"`
	OAuth     bool           `db:"oauth"`
	Email     string         `db:"email"`
	Password  sql.NullString `db:"password"`
	CreatedAt time.Time      `db:"created_at"`
}
