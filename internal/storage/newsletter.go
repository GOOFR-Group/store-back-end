package storage

import (
	"github.com/gocraft/dbr/v2"
)

const NewsletterTable = "newsletter"

// tgcon - used to generate constants for each field's tag
type Newsletter struct {
	Email string `db:"email"`
}

func CreateNewsletter(t Transaction, model Newsletter) error {
	_, err := t.InsertInto(NewsletterTable).
		Columns(NewsletterEmailDb).
		Record(model).
		Exec()

	return err
}

func ReadNewsletters(t Transaction) (objects []Newsletter, err error) {
	_, err = t.Select("*").
		From(NewsletterTable).
		Load(&objects)

	return
}

func ReadNewsletterByID(t Transaction, email string) (object Newsletter, ok bool, err error) {
	err = t.Select("*").
		From(NewsletterTable).
		Where(NewsletterEmailDb+" = ?", email).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
func DeleteNewsletterByID(t Transaction, email string) (ok bool, err error) {
	_, err = t.DeleteFrom(NewsletterTable).
		Where(NewsletterEmailDb+" = ?", email).
		Exec()

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}
