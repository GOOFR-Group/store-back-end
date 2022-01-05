package storage

import (
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
