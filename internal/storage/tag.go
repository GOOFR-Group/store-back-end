package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const TagTable = "tag"

// tgcon - used to generate constants for each field's tag
type Tag struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func CreateNewTag(t Transaction, model Tag) error {
	_, err := t.InsertInto(TagTable).
		Columns(TagIDDb, TagNameDb).
		Record(model).
		Exec()

	return err
}

func GetAllTags(t Transaction) (objects []Tag, err error) {
	_, err = t.Select("*").
		From(TagTable).
		Load(&objects)

	return
}

func GetTagByID(t Transaction, id uuid.UUID) (object Tag, ok bool, err error) {
	err = t.Select("*").
		From(TagTable).
		Where(TagIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateTagByID(t Transaction, model Tag) (ok bool, err error) {
	_, err = t.Update(TagTable).
		SetMap(map[string]interface{}{
			TagNameDb: model.Name,
		}).
		Where(TagIDDb+" = ?", model.ID).
		Exec()

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func DeleteTagByID(t Transaction, id uuid.UUID) (ok bool, err error) {
	_, err = t.DeleteFrom(TagTable).
		Where(TagIDDb+" = ?", id).
		Exec()

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
