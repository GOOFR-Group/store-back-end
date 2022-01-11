package core

import "errors"

const (
	ErrGeneratingUUID = `failed to generate new UUID: %s`
)

// general
var (
	ErrObjectNotFound       = errors.New("failed to find object")
	ErrObjectAlreadyCreated = errors.New("object already created")
)

// newsletter
var (
	ErrInvalidEmail = errors.New("invalid email")
)
