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

func CreateClientSearchHistory(t Transaction, model ClientSearchHistory) error {
	_, err := t.InsertInto(ClientSearchHistoryTable).
		Columns(ClientSearchHistoryIDDb, ClientSearchHistoryIDGameDb, ClientSearchHistoryIDClientDb, ClientSearchHistoryDateTimeDb).
		Record(model).
		Exec()

	return err
}

func ReadClientSearchHistoryByClientID(t Transaction, id uuid.UUID) (objects []ClientSearchHistory, err error) {
	_, err = t.Select("*").
		From(ClientSearchHistoryTable).
		Where(ClientSearchHistoryIDClientDb+" = ?", id).
		Load(&objects)

	return
}
