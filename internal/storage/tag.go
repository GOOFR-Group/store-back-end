package storage

import (
	"github.com/google/uuid"
)

const TagTable = "tag"

// tgcon - used to generate constants for each field's tag
type Tag struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
