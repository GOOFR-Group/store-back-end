package storage

import (
	"database/sql"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const AccessTable = "access"

// tgcon - used to generate constants for each field's tag
type Access struct {
	ID       uuid.UUID      `db:"id"`
	IDClient uuid.UUID      `db:"id_client"`
	OAuth    bool           `db:"oauth"`
	Email    string         `db:"email"`
	Password sql.NullString `db:"password"`
}

func CreateAccess(t Transaction, model Access) error {
	_, err := t.InsertInto(AccessTable).
		Columns(AccessIDDb, AccessIDClientDb, AccessOAuthDb, AccessEmailDb, AccessPasswordDb).
		Record(model).
		Exec()

	return err
}

func ReadAccessByClientID(t Transaction, id uuid.UUID) (object Access, ok bool, err error) {
	err = t.Select("*").
		From(AccessTable).
		Where(AccessIDClientDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadAccessByEmail(t Transaction, email string) (object Access, ok bool, err error) {
	err = t.Select("*").
		From(AccessTable).
		Where(AccessEmailDb+" = ?", email).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadAccessByEmailNotFromClientID(t Transaction, email string, id uuid.UUID) (object Access, ok bool, err error) {
	err = t.Select("*").
		From(AccessTable).
		Where(AccessEmailDb+" = ?", email).
		Where(AccessIDClientDb+" != ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateAccessByClientID(t Transaction, model Access) error {
	_, err := t.Update(AccessTable).
		SetMap(map[string]interface{}{
			AccessOAuthDb:    model.OAuth,
			AccessEmailDb:    model.Email,
			AccessPasswordDb: model.Password,
		}).
		Where(AccessIDClientDb+" = ?", model.IDClient).
		Exec()

	return err
}
