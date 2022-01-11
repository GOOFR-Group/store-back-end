package core

import "errors"

const (
	ErrGeneratingUUID = `failed to generate new UUID: %s`
)

var (
	ErrObjectNotFound = errors.New("failed to find object")
)
