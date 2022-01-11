package core

import (
	"database/sql"
	"fmt"
	"unicode"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// PostLogin log in to a client's account
func PostLogin(req oapi.PostLoginJSONRequestBody) (oapi.ClientSchema, error) {
	var access storage.Access
	var object storage.Client

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool
		var err error

		if access, ok, err = storage.ReadAccessByEmail(tx, req.Email); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if object, ok, err = storage.ReadClientByID(tx, access.IDClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		return nil
	}); err != nil {
		return oapi.ClientSchema{}, err
	}

	if req.Password == nil && !access.OAuth {
		return oapi.ClientSchema{}, ErrPasswordRequired
	}

	if req.Password != nil && !access.OAuth {
		if !access.Password.Valid {
			return oapi.ClientSchema{}, ErrIncorrectPassword
		}
		if !checkPasswordHash(*req.Password, access.Password.String) {
			return oapi.ClientSchema{}, ErrIncorrectPassword
		}
	}

	if !object.Active {
		return oapi.ClientSchema{}, ErrClientInactive
	}

	return getClientFromModel(object), nil
}

// PostRegister registers a client
func PostRegister(req oapi.PostRegisterJSONRequestBody) (oapi.ClientSchema, error) {
	if !validEmail(req.Access.Email) {
		return oapi.ClientSchema{}, ErrInvalidEmail
	}

	if req.Access.Password == nil && !req.Access.Oauth {
		return oapi.ClientSchema{}, ErrPasswordRequired
	}

	var password string
	var passwordOK bool
	if req.Access.Password != nil && !req.Access.Oauth {
		password = *req.Access.Password

		if !validPassword(password) {
			return oapi.ClientSchema{}, ErrInvalidPassword
		}

		var err error
		if password, err = hashPassword(password); err != nil {
			return oapi.ClientSchema{}, err
		}

		passwordOK = true
	}

	var id uuid.UUID
	var err error

	if id, err = uuid.NewRandom(); err != nil {
		return oapi.ClientSchema{}, fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	var object storage.Client

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadAccessByEmail(tx, req.Access.Email); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		if err = storage.CreateClient(tx, storage.Client{
			ID:          id,
			Name:        req.Client.Name,
			Surname:     req.Client.Surname,
			Picture:     req.Client.Picture,
			Birthdate:   req.Client.Birthdate.Time,
			PhoneNumber: req.Client.PhoneNumber,
			VatID:       req.Client.VatId,
			Active:      req.Client.Active,
		}); err != nil {
			return err
		}

		if object, ok, err = storage.ReadClientByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if id, err = uuid.NewRandom(); err != nil {
			return fmt.Errorf(ErrGeneratingUUID, err.Error())
		}

		if err = storage.CreateAccess(tx, storage.Access{
			ID:       id,
			IDClient: object.ID,
			OAuth:    req.Access.Oauth,
			Email:    req.Access.Email,
			Password: sql.NullString{
				String: password,
				Valid:  passwordOK,
			},
		}); err != nil {
			return err
		}

		if id, err = uuid.NewRandom(); err != nil {
			return fmt.Errorf(ErrGeneratingUUID, err.Error())
		}

		var doorNumber string
		var doorNumberOK bool
		if req.Address.DoorNumber != nil {
			doorNumber = *req.Address.DoorNumber
			doorNumberOK = true
		}

		if err = storage.CreateAddress(tx, storage.Address{
			ID:       id,
			IDClient: object.ID,
			Street:   req.Address.Street,
			DoorNumber: sql.NullString{
				String: doorNumber,
				Valid:  doorNumberOK,
			},
			ZipCode: req.Address.ZipCode,
			City:    req.Address.City,
			Country: req.Address.Country,
		}); err != nil {
			return err
		}

		if id, err = uuid.NewRandom(); err != nil {
			return fmt.Errorf(ErrGeneratingUUID, err.Error())
		}

		if err = storage.CreateWallet(tx, storage.Wallet{
			ID:       id,
			IDClient: object.ID,
			Balance:  0,
			Coin:     req.Wallet.Coin,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.ClientSchema{}, err
	}

	return getClientFromModel(object), nil
}

// PutAccess updates a client access
func PutAccess(params oapi.PutAccessParams, req oapi.PutAccessJSONRequestBody) error {
	if !validEmail(req.Email) {
		return ErrInvalidEmail
	}

	if req.Password == nil && !req.Oauth {
		return ErrPasswordRequired
	}

	var password string
	var passwordOK bool
	if req.Password != nil && !req.Oauth {
		password = *req.Password

		if !validPassword(password) {
			return ErrInvalidPassword
		}

		var err error
		if password, err = hashPassword(password); err != nil {
			return err
		}

		passwordOK = true
	}

	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.ClientID); err != nil {
		return err
	}

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var client storage.Client
		var ok bool

		if client, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if _, ok, err = storage.ReadAccessByEmailNotFromClientID(tx, req.Email, client.ID); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		if _, ok, err = storage.ReadAccessByClientID(tx, client.ID); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.UpdateAccessByClientID(tx, storage.Access{
			IDClient: client.ID,
			OAuth:    req.Oauth,
			Email:    req.Email,
			Password: sql.NullString{
				String: password,
				Valid:  passwordOK,
			},
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func validPassword(password string) bool {
	digits := 0
	specialChars := 0
	for _, r := range password {
		if unicode.IsDigit(r) {
			digits++
			continue
		}
		if !unicode.IsLetter(r) {
			specialChars++
			continue
		}
	}
	return len(password) >= 6 && digits >= 1 && specialChars >= 1
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
