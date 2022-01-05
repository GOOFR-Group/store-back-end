package storage

import (
	"github.com/google/uuid"
)

const NewsletterTable = "newsletter"

// tgcon - used to generate constants for each field's tag
type Newsletter struct {
	Email uuid.UUID `db:"email"`
}
