package storage

import (
	"strings"
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

func ReadGamesWithDifferentID(t Transaction, ids []uuid.UUID, limit int64) (objects []Game, err error) {
	idsString := make([]string, len(ids))
	for i, id := range ids {
		idsString[i] = "'" + id.String() + "'"
	}

	if len(ids) == 0 {
		_, err = t.Select("*").
			From(GameTable).
			Limit(uint64(limit)).
			Load(&objects)

		return
	}

	_, err = t.Select("*").
		From(GameTable).
		Where(GameIDDb + " NOT IN (" + strings.Join(idsString, ", ") + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesWithDiscount(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		Where(GameDiscountDb+" > ?", 0).
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesByNameLike(t Transaction, like string, limit int64) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		Where(GameNameDb+" LIKE '%?%'", like).
		Limit(uint64(limit)).
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

func ReadGamesAndAverageOrderByAvgReviewDesc(t Transaction, limit int64) (objects []Game, averages []float64, err error) {
	if _, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(ReviewTable, GameTable+"."+GameIDDb+" = "+ReviewTable+"."+ReviewIDGameDb).
		Where(ReviewTable + "." + ReviewStarsDb + " IS NOT NULL").
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("AVG(" + ReviewTable + "." + ReviewStarsDb + ")").
		Limit(uint64(limit)).
		Load(&objects); err != nil {
		return
	}

	_, err = t.Select("AVG("+ReviewTable+"."+ReviewStarsDb+")").
		From(GameTable).
		Join(ReviewTable, GameTable+"."+GameIDDb+" = "+ReviewTable+"."+ReviewIDGameDb).
		Where(ReviewTable + "." + ReviewStarsDb + " IS NOT NULL").
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("AVG(" + ReviewTable + "." + ReviewStarsDb + ")").
		Limit(uint64(limit)).
		Load(&averages)

	return
}

func ReadGamesOrderByReleaseDateDescAndByAvgReviewDesc(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(ReviewTable, GameTable+"."+GameIDDb+" = "+ReviewTable+"."+ReviewIDGameDb).
		Where(ReviewTable + "." + ReviewStarsDb + " IS NOT NULL").
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc(GameTable + "." + GameReleaseDateDb).
		OrderDesc("AVG(" + ReviewTable + "." + ReviewStarsDb + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesOrderByReleaseDateDesc(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select("*").
		From(GameTable).
		OrderDesc(GameReleaseDateDb).
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesOrderByReleaseDateDescFilteredByTag(t Transaction, tags []uuid.UUID, limit int64) (objects []Game, err error) {
	if len(tags) == 0 {
		return
	}

	_, err = t.Select("*").
		From(GameTable).
		Where(GameIDDb+" IN (?)", readGamesIDFilteredByTag(t, tags)).
		OrderDesc(GameReleaseDateDb).
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesMostPurchasedOrderByAvgReviewDesc(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(ReviewTable, GameTable+"."+GameIDDb+" = "+ReviewTable+"."+ReviewIDGameDb).
		Join(GameLibraryTable, GameTable+"."+GameIDDb+" = "+GameLibraryTable+"."+GameLibraryIDGameDb).
		Where(ReviewTable + "." + ReviewStarsDb + " IS NOT NULL").
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("COUNT(" + GameLibraryTable + "." + GameLibraryIDGameDb + ")").
		OrderDesc("AVG(" + ReviewTable + "." + ReviewStarsDb + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesMostPurchased(t Transaction, limit int64) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(GameLibraryTable, GameTable+"."+GameIDDb+" = "+GameLibraryTable+"."+GameLibraryIDGameDb).
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("COUNT(" + GameLibraryTable + "." + GameLibraryIDGameDb + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesMostPurchasedFilteredByTag(t Transaction, tags []uuid.UUID, limit int64) (objects []Game, err error) {
	if len(tags) == 0 {
		return
	}

	_, err = t.Select(GameTable+".*").
		From(GameTable).
		Join(GameLibraryTable, GameTable+"."+GameIDDb+" = "+GameLibraryTable+"."+GameLibraryIDGameDb).
		Where(GameTable+"."+GameIDDb+" IN (?)", readGamesIDFilteredByTag(t, tags)).
		GroupBy(GameTable + "." + GameIDDb).
		OrderDesc("COUNT(" + GameLibraryTable + "." + GameLibraryIDGameDb + ")").
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadGamesRecommendedByClientID(t Transaction, id uuid.UUID, limit int64) (objects []Game, err error) {
	var tags []Tag
	if tags, err = ReadTagsByClientID(t, id); err != nil {
		return
	}

	tagsID := make([]uuid.UUID, len(tags))
	for i, tag := range tags {
		tagsID[i] = tag.ID
	}

	if len(tagsID) == 0 {
		return
	}

	_, err = t.Select("*").
		From(GameTable).
		Where(GameIDDb+" IN (?)", readGamesIDFilteredByTag(t, tagsID)).
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

func readGamesIDFilteredByTag(t Transaction, tags []uuid.UUID) *dbr.SelectStmt {
	if len(tags) == 0 {
		return t.Select()
	}

	tagsString := make([]string, len(tags))
	for i, t := range tags {
		tagsString[i] = "'" + t.String() + "'"
	}

	return t.Select("DISTINCT "+GameTable+"."+GameIDDb).
		From(GameTable).
		Join(TagGameTable, GameTable+"."+GameIDDb+" = "+TagGameTable+"."+TagGameIDGameDb).
		Join(TagTable, TagGameTable+"."+TagGameIDTagDb+" = "+TagTable+"."+TagIDDb).
		Where(TagTable + "." + TagIDDb + " IN (" + strings.Join(tagsString, ", ") + ")")
}
