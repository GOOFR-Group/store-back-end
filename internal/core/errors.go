package core

import "errors"

const (
	ErrGeneratingUUID = `failed to generate new UUID: %s`
)

// general
var (
	ErrObjectNotFound       = errors.New("failed to find object")
	ErrObjectAlreadyCreated = errors.New("object already created")
	ErrInvalidEmail         = errors.New("invalid email")
)

// specific
var (
	ErrGameNotFound      = errors.New("failed to find game")
	ErrPublisherNotFound = errors.New("failed to find publisher")
	ErrTagNotFound       = errors.New("failed to find tag")
)
