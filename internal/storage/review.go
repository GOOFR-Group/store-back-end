package storage

import (
	"database/sql"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const ReviewTable = "review"

// tgcon - used to generate constants for each field's tag
type Review struct {
	ID       uuid.UUID      `db:"id"`
	IDGame   uuid.UUID      `db:"id_game"`
	IDClient uuid.UUID      `db:"id_client"`
	Stars    sql.NullInt64  `db:"stars"`
	Review   sql.NullString `db:"review"`
}

func CreateReview(t Transaction, model Review) error {
	_, err := t.InsertInto(ReviewTable).
		Columns(ReviewIDDb, ReviewIDGameDb, ReviewIDClientDb, ReviewStarsDb, ReviewReviewDb).
		Record(model).
		Exec()

	return err
}

func ReadReviewsByGameID(t Transaction, id uuid.UUID) (objects []Review, err error) {
	_, err = t.Select("*").
		From(ReviewTable).
		Where(ReviewIDGameDb+" = ?", id).
		Load(&objects)

	return
}

func ReadReviewByID(t Transaction, id uuid.UUID) (object Review, ok bool, err error) {
	err = t.Select("*").
		From(ReviewTable).
		Where(ReviewIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadReviewByGameIDAndClientID(t Transaction, gameID, clientID uuid.UUID) (object Review, ok bool, err error) {
	err = t.Select("*").
		From(ReviewTable).
		Where(ReviewIDGameDb+" = ?", gameID).
		Where(ReviewIDClientDb+" = ?", clientID).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateReviewByID(t Transaction, model Review) error {
	_, err := t.Update(ReviewTable).
		SetMap(map[string]interface{}{
			ReviewStarsDb:  model.Stars,
			ReviewReviewDb: model.Review,
		}).
		Where(ReviewIDDb+" = ?", model.ID).
		Exec()

	return err
}

func DeleteReviewByID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(ReviewTable).
		Where(ReviewIDDb+" = ?", id).
		Exec()

	return err
}
