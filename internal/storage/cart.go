package storage

import (
	"github.com/google/uuid"
)

const CartTable = "cart"

// tgcon - used to generate constants for each field's tag
type Cart struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}
