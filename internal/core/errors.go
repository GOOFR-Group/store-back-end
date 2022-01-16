package core

import "errors"

const (
	ErrGeneratingUUID = "failed to generate new UUID: %s"
)

// general
var (
	ErrObjectNotFound       = errors.New("failed to find object")
	ErrObjectAlreadyCreated = errors.New("object already created")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrPasswordRequired     = errors.New("password required")
	ErrIncorrectPassword    = errors.New("incorrect password")
	ErrInvalidAmount        = errors.New("invalid amount")
)

// specific
var (
	ErrGameNotFound          = errors.New("failed to find game")
	ErrPublisherNotFound     = errors.New("failed to find publisher")
	ErrClientNotFound        = errors.New("failed to find client")
	ErrTagNotFound           = errors.New("failed to find tag")
	ErrInvoiceHeaderNotFound = errors.New("failed to find invoice header")
	ErrClientInactive        = errors.New("inactive client")
	ErrGameAlreadyBought     = errors.New("game already bought")
)
