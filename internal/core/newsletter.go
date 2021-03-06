package core

import (
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/conf"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

const (
	smtpSubject = "Subject: %s\r\n\r\n"
	smtpMIME    = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// PostNewsletter adds an email to the newsletter list
func PostNewsletter(params oapi.PostNewsletterParams) error {
	if !validEmail(params.Email) {
		return ErrInvalidEmail
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool
		var err error

		if _, ok, err = storage.ReadNewsletterByID(tx, params.Email); err != nil {
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
func GetNewsletter(params oapi.GetNewsletterParams) ([]oapi.NewsletterSchema, error) {
	var objects []storage.Newsletter

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if params.Email == nil {
			if objects, err = storage.ReadNewsletters(tx); err != nil {
				return err
			}
		} else {
			var object storage.Newsletter
			var ok bool

			if object, ok, err = storage.ReadNewsletterByID(tx, *params.Email); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			objects = append(objects, object)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getNewslettersFromModel(objects), nil
}

// DeleteNewsletter removes an email from the newsletter list
func DeleteNewsletter(params oapi.DeleteNewsletterParams) (oapi.NewsletterSchema, error) {
	var object storage.Newsletter

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool
		var err error

		if object, ok, err = storage.ReadNewsletterByID(tx, params.Email); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteNewsletterByID(tx, params.Email); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.NewsletterSchema{}, err
	}

	return getNewsletterFromModel(object), nil
}

// PostSendNewsletter sends a newsletter to all registered emails
func PostSendNewsletter(req oapi.PostSendNewsletterJSONRequestBody) error {
	var objects []storage.Newsletter

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error
		objects, err = storage.ReadNewsletters(tx)

		return err
	}); err != nil {
		return err
	}

	var to []string
	for _, e := range objects {
		if !validEmail(e.Email) {
			continue
		}
		to = append(to, e.Email)
	}

	title := "GOOFR Store - " + req.Title

	body := `<html> <body style="background-color: #0D1B2A; font-family: sans-serif; padding-top: 20px; padding-bottom: 20px;">`
	body += `<h1 style="text-align: center; color: #778DA9;">` + title + `</h1> <hr style="border-color: #778DA9;"> <br>`

	var idPublisher uuid.UUID
	var publisher storage.Publisher
	var err error

	for _, g := range req.Games {
		if idPublisher, err = uuid.Parse(g.IdPublisher); err != nil {
			return err
		}

		if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
			var ok bool

			if publisher, ok, err = storage.ReadPublisherByID(tx, idPublisher); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			return nil
		}); err != nil {
			return err
		}

		body += `<div style="margin: 0px 50px; margin-bottom: 15px;">`
		body += fmt.Sprintf(`<img style="width: auto; height: 135px;" src="%s">`, g.CoverImage)
		body += `<div style="padding-top: 10px; float: right; color: #E0E1DD;"> <table> `
		body += fmt.Sprintf(`<tr> <th style="text-align: right;">%s</th> <td style="padding-left: 15px;">%s</td> </tr>`, "Name", g.Name)
		body += fmt.Sprintf(`<tr> <th style="text-align: right;">%s</th> <td style="padding-left: 15px;">%s</td> </tr>`, "Publisher", publisher.Name)
		body += fmt.Sprintf(`<tr> <th style="text-align: right;">%s</th> <td style="padding-left: 15px;">???%.2f</td> </tr>`, "Price", g.Price)
		body += fmt.Sprintf(`<tr> <th style="text-align: right;">%s</th> <td style="padding-left: 15px;">%.2f%%</td> </tr>`, "Discount", g.Discount*100)
		body += fmt.Sprintf(`<tr> <th style="text-align: right;">%s</th> <td style="padding-left: 15px;">%s</td> </tr>`, "Release Date", g.ReleaseDate.Format(timeLayout))
		body += `</table> </div> </div>`
	}

	body += `</body> </html>`

	message := []byte(fmt.Sprintf(smtpSubject, title) + smtpMIME + body)
	return smtp.SendMail(conf.SMTPAddress(), conf.SMTPAuthentication(), conf.SMTPEmailAddress(), to, message)
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
