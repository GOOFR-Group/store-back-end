package storage

import (
	"time"

	"github.com/google/uuid"
)

const ClientTable = "client"

// tgcon - used to generate constants for each field's tag
type Round struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Surname     string    `db:"surname"`
	Picture     string    `db:"picture"`
	Birthdate   time.Time `db:"birthdate"`
	PhoneNumber string    `db:"phone_number"`
	VatID       int64     `db:"vat_id"`
	Active      bool      `db:"active"`
}
