package storage

import (
	"github.com/google/uuid"
)

const WalletTable = "wallet"

// tgcon - used to generate constants for each field's tag
type Wallet struct {
	ID       uuid.UUID `db:"id"`
	IDClient uuid.UUID `db:"id_client"`
	Balance  float64   `db:"balance"`
	Coin     byte      `db:"coin"`
}
