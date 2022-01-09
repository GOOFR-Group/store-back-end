package storage

import (
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/utils/mathf"
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

func CreateGame(t Transaction, model Game) error {
	model.Discount = mathf.Clamp(model.Discount, 0, 1)
	_, err := t.InsertInto(GameTable).
		Columns(GameIDDb, GameIDPublisherDb, GameNameDb, GamePriceDb, GameDiscountDb, GameStateDb, GameCoverImageDb, GameReleaseDateDb, GameDescriptionDb, GameDownloadLinkDb).
		Record(model).
		Exec()

	return err
}

func ReadGames(t Transaction) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		Load(&objects)

	return
}

func ReadGamesOrderByAvgReviewDesc(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(ReviewTable, GameTable+"."+GameIDDb+" = "+ReviewTable+"."+ReviewIDGameDb).
		Where(ReviewTable + "." + ReviewStarsDb + " IS NOT NULL").
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("AVG(" + ReviewTable + "." + ReviewStarsDb + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesByPublisherID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		Where(GameIDPublisherDb+" = ?", id).
		Load(&objects)

	return
}

func ReadGameByID(t Transaction, id uuid.UUID) (object Game, ok bool, err error) {
	err = t.Select("*").
		From(GameTable).
		Where(GameIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateGameByID(t Transaction, model Game) error {
	model.Discount = mathf.Clamp(model.Discount, 0, 1)
	_, err := t.Update(GameTable).
		SetMap(map[string]interface{}{
			GameIDPublisherDb:  model.IDPublisher,
			GameNameDb:         model.Name,
			GamePriceDb:        model.Price,
			GameDiscountDb:     model.Discount,
			GameStateDb:        model.State,
			GameCoverImageDb:   model.CoverImage,
			GameReleaseDateDb:  model.ReleaseDate,
			GameDescriptionDb:  model.Description,
			GameDownloadLinkDb: model.DownloadLink,
		}).
		Where(GameIDDb+" = ?", model.ID).
		Exec()

	return err
}

func DeleteGameByID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(GameTable).
		Where(GameIDDb+" = ?", id).
		Exec()

	return err
}
