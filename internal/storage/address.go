package storage

import (
	"database/sql"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const AddressTable = "address"

// tgcon - used to generate constants for each field's tag
type Address struct {
	ID         uuid.UUID      `db:"id"`
	IDClient   uuid.UUID      `db:"id_client"`
	Street     string         `db:"street"`
	DoorNumber sql.NullString `db:"door_number"`
	ZipCode    string         `db:"zip_code"`
	City       string         `db:"city"`
	Country    string         `db:"country"`
}

func CreateAddress(t Transaction, model Address) error {
	_, err := t.InsertInto(AddressTable).
		Columns(AddressIDDb, AddressIDClientDb, AddressStreetDb, AddressDoorNumberDb, AddressZipCodeDb, AddressCityDb, AddressCountryDb).
		Record(model).
		Exec()

	return err
}

func ReadAddressByID(t Transaction, id uuid.UUID) (object Address, ok bool, err error) {
	err = t.Select("*").
		From(AddressTable).
		Where(AddressIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadAddressByClientID(t Transaction, id uuid.UUID) (object Address, ok bool, err error) {
	err = t.Select("*").
		From(AddressTable).
		Where(AddressIDClientDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateAddressByID(t Transaction, model Address) error {
	_, err := t.Update(AddressTable).
		SetMap(map[string]interface{}{
			AddressStreetDb:     model.Street,
			AddressDoorNumberDb: model.DoorNumber,
			AddressZipCodeDb:    model.ZipCode,
			AddressCityDb:       model.City,
			AddressCountryDb:    model.Country,
		}).
		Where(AddressIDDb+" = ?", model.ID).
		Exec()

	return err
}
