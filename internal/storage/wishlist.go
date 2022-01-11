package storage

import (
	"github.com/google/uuid"
)

const WishlistTable = "wishlist"

// tgcon - used to generate constants for each field's tag
type Wishlist struct {
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
}
