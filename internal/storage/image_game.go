package storage

import (
	"github.com/google/uuid"
)

const ImageGameTable = "image_game"

// tgcon - used to generate constants for each field's tag
type ImageGame struct {
	ID     uuid.UUID `db:"id"`
	IDGame uuid.UUID `db:"id_game"`
	Image  string    `db:"image"`
}
