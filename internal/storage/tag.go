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

func CreateTag(t Transaction, model Tag) error {
	_, err := t.InsertInto(TagTable).
		Columns(TagIDDb, TagNameDb).
		Record(model).
		Exec()

	return err
}

func ReadTags(t Transaction) (objects []Tag, err error) {
	_, err = t.Select("*").
		From(TagTable).
		Load(&objects)

	return
}

func ReadTagsByNameLike(t Transaction, like string, limit int64) (objects []Tag, err error) {
	_, err = t.Select("*").
		From(TagTable).
		Where(TagNameDb+" LIKE '%?%'", like).
		Limit(uint64(limit)).
		Load(&objects)

	return
}

func ReadTagsByGameID(t Transaction, id uuid.UUID) (objects []Tag, err error) {
	_, err = t.Select(TagTable+".*").
		From(TagTable).
		Join(TagGameTable, TagTable+"."+TagIDDb+" = "+TagGameTable+"."+TagGameIDTagDb).
		Where(TagGameTable+"."+TagGameIDGameDb+" = ?", id).
		Load(&objects)

	return
}

func ReadTagsByClientID(t Transaction, id uuid.UUID) (objects []Tag, err error) {
	_, err = t.Select("DISTINCT "+TagTable+".*").
		From(TagTable).
		FullJoin(TagGameTable, TagTable+"."+TagIDDb+" = "+TagGameTable+"."+TagGameIDTagDb).
		FullJoin(GameTable, TagGameTable+"."+TagGameIDGameDb+" = "+GameTable+"."+GameIDDb).
		FullJoin(ClientSearchHistoryTable, GameTable+"."+GameIDDb+" = "+ClientSearchHistoryTable+"."+ClientSearchHistoryIDGameDb).
		FullJoin(GameLibraryTable, GameTable+"."+GameIDDb+" = "+GameLibraryTable+"."+GameLibraryIDGameDb).
		Where("("+ClientSearchHistoryTable+"."+ClientSearchHistoryIDClientDb+" = ? OR "+GameLibraryTable+"."+GameLibraryIDClientDb+" = ?) AND ("+TagTable+"."+TagIDDb+" IS NOT NULL AND "+TagTable+"."+TagNameDb+" IS NOT NULL)", id, id).
		Load(&objects)

	return
}

func ReadTagByID(t Transaction, id uuid.UUID) (object Tag, ok bool, err error) {
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

func UpdateTagByID(t Transaction, model Tag) error {
	_, err := t.Update(TagTable).
		SetMap(map[string]interface{}{
			TagNameDb: model.Name,
		}).
		Where(TagIDDb+" = ?", model.ID).
		Exec()

	return err
}

func DeleteTagByID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(TagTable).
		Where(TagIDDb+" = ?", id).
		Exec()

	return err
}
