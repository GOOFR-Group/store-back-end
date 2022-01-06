package storage

import (
	"time"

	"github.com/google/uuid"
)

const GameTable = "game"

// tgcon - used to generate constants for each field's tag
type Game struct {
	ID           uuid.UUID `db:"id"`
	IDPublisher  uuid.UUID `db:"id_publisher"`
	Name         string    `db:"name"`
	Price        float64   `db:"price"`
	Discount     float64   `db:"discount"`
	State        StateGame `db:"state"`
	CoverImage   string    `db:"cover_image"`
	ReleaseDate  time.Time `db:"release_date"`
	Description  string    `db:"description"`
	DownloadLink string    `db:"download_link"`
}

func ReadGamesByPublisherID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		Where(GameIDPublisherDb+" = ?", id).
		Load(&objects)

	return
}
