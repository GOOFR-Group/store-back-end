package handlers

// general
const (
	// internal
	ErrWritingResponse = "failed to write response"
	ErrInternalServer  = "internal server error"

	// response
	ErrParsingRequest = `error parsing request - %s`
)

// tag
const (
	ErrTagNotFound = `tag not found with ID: %s`
)
