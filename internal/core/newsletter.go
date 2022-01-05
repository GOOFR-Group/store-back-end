package core

import (
	"net/mail"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
	"github.com/gocraft/dbr/v2"
)

// PostNewsletter adds an email to the newsletter list
func PostNewsletter(params oapi.PostNewsletterParams) error {
	if !validEmail(params.Email) {
		return ErrInvalidEmail
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		_, ok, err := storage.ReadNewsletterByID(tx, params.Email)
		if err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		return storage.CreateNewsletter(tx, storage.Newsletter{
			Email: params.Email,
		})
	}); err != nil {
		return err
	}

	return nil
}

// GetNewsletter gets the list of email subscribed to the newsletter
func GetNewsletter() ([]oapi.NewsletterSchema, error) {
	var objects []storage.Newsletter

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error
		objects, err = storage.ReadAllNewsletters(tx)

		return err
	}); err != nil {
		return nil, err
	}

	return getNewslettersFromModel(objects), nil
}

// DeleteNewsletter removes an email from the newsletter list
func DeleteNewsletter(params oapi.DeleteNewsletterParams) (oapi.NewsletterSchema, error) {
	var object storage.Newsletter

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error
		var ok bool

		object, ok, err = storage.ReadNewsletterByID(tx, params.Email)
		if err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		ok, err = storage.DeleteNewsletterByID(tx, params.Email)
		if err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.NewsletterSchema{}, err
	}

	return getNewsletterFromModel(object), nil
}

// PostSendNewsletter sends a newsletter to all registered emails
func PostSendNewsletter(req oapi.PostSendNewsletterJSONRequestBody) error {

	return nil
}

func getNewsletterFromModel(model storage.Newsletter) oapi.NewsletterSchema {
	return oapi.NewsletterSchema{
		Email: model.Email,
	}
}

func getNewslettersFromModel(model []storage.Newsletter) []oapi.NewsletterSchema {
	array := make([]oapi.NewsletterSchema, len(model))
	for i, m := range model {
		array[i] = getNewsletterFromModel(m)
	}
	return array
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
