package storage

import (
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const ClientTable = "client"

// tgcon - used to generate constants for each field's tag
type Client struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Surname     string    `db:"surname"`
	Picture     string    `db:"picture"`
	Birthdate   time.Time `db:"birthdate"`
	PhoneNumber string    `db:"phone_number"`
	VatID       int64     `db:"vat_id"`
	Active      bool      `db:"active"`
}

func CreateClient(t Transaction, model Client) error {
	_, err := t.InsertInto(ClientTable).
		Columns(ClientIDDb, ClientNameDb, ClientSurnameDb, ClientPictureDb, ClientBirthdateDb, ClientPhoneNumberDb, ClientVatIDDb, ClientActiveDb).
		Record(model).
		Exec()

	return err
}

func ReadClients(t Transaction) (objects []Client, err error) {
	_, err = t.Select("*").
		From(ClientTable).
		Load(&objects)

	return
}

func ReadClientByID(t Transaction, id uuid.UUID) (object Client, ok bool, err error) {
	err = t.Select("*").
		From(ClientTable).
		Where(ClientIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateClientByID(t Transaction, model Client) error {
	_, err := t.Update(ClientTable).
		SetMap(map[string]interface{}{
			ClientNameDb:        model.Name,
			ClientSurnameDb:     model.Surname,
			ClientPictureDb:     model.Picture,
			ClientBirthdateDb:   model.Birthdate,
			ClientPhoneNumberDb: model.PhoneNumber,
			ClientVatIDDb:       model.VatID,
			ClientActiveDb:      model.Active,
		}).
		Where(ClientIDDb+" = ?", model.ID).
		Exec()

	return err
}

func DeleteClientByID(t Transaction, id uuid.UUID) error {
	_, err := t.DeleteFrom(ClientTable).
		Where(ClientIDDb+" = ?", id).
		Exec()

	return err
}
