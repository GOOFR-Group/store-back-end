package storage

import (
	"time"

	"github.com/google/uuid"
)

const ClientSearchHistoryTable = "client_search_history"

// tgcon - used to generate constants for each field's tag
type ClientSearchHistory struct {
	ID       uuid.UUID `db:"id"`
	IDGame   uuid.UUID `db:"id_game"`
	IDClient uuid.UUID `db:"id_client"`
	DateTime time.Time `db:"date_time"`
}
