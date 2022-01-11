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

func DeleteWishlistByID(t Transaction, gameID, clientID uuid.UUID) error {
	_, err := t.DeleteFrom(WishlistTable).
		Where(WishlistIDGameDb+" = ?", gameID).
		Where(WishlistIDClientDb+" = ?", clientID).
		Exec()

	return err
}
