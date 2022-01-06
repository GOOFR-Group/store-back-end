package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const PublisherTable = "publisher"

// tgcon - used to generate constants for each field's tag
type Publisher struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	CoverImage  string    `db:"cover_image"`
	PhoneNumber string    `db:"phone_number"`
	Email       string    `db:"email"`
}

func CreatePublisher(t Transaction, model Publisher) error {
	_, err := t.InsertInto(PublisherTable).
		Columns(PublisherIDDb, PublisherNameDb, PublisherCoverImageDb, PublisherPhoneNumberDb, PublisherEmailDb).
		Record(model).
		Exec()

	return err
}

func ReadAllPublishers(t Transaction) (objects []Publisher, err error) {
	_, err = t.Select("*").
		From(PublisherTable).
		Load(&objects)

	return
}

func ReadPublisherByID(t Transaction, id uuid.UUID) (object Publisher, ok bool, err error) {
	err = t.Select("*").
		From(PublisherTable).
		Where(PublisherIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
func ReadPublisherGamesByID(t Transaction, id uuid.UUID) (objects []Game, err error) {
	_, err = t.Select(GameTable+".*").
		From(PublisherTable).
		Join(GameTable, PublisherTable+"."+PublisherIDDb+" = "+GameTable+"."+GameIDPublisherDb).
		Load(&objects)

	return
}

func UpdatePublisherByID(t Transaction, model Publisher) error {
	_, err := t.Update(PublisherTable).
		SetMap(map[string]interface{}{
			PublisherNameDb:        model.Name,
			PublisherCoverImageDb:  model.CoverImage,
			PublisherPhoneNumberDb: model.PhoneNumber,
			PublisherEmailDb:       model.Email,
		}).
		Where(PublisherIDDb+" = ?", model.ID).
		Exec()

	return err
}

func DeletePublisherByID(t Transaction, id uuid.UUID) (ok bool, err error) {
	_, err = t.DeleteFrom(PublisherTable).
		Where(PublisherIDDb+" = ?", id).
		Exec()

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
