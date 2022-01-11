package storage

import (
	"github.com/google/uuid"
)

const TagGameTable = "tag_game"

// tgcon - used to generate constants for each field's tag
type TagGame struct {
	IDTag  uuid.UUID `db:"id_tag"`
	IDGame uuid.UUID `db:"id_game"`
}
