package storage

import (
	"database/sql"

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
