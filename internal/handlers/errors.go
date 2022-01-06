package handlers

// general
const (
	// internal
	ErrWritingResponse = "failed to write response"
	ErrInternalServer  = "internal server error"

	// response
	ErrParsingRequest = `error parsing request: %s`
)

// publisher
const (
	ErrPublisherNotFound = `publisher not found with ID: %s`
)

// tag
const (
	ErrTagNotFound = `tag not found with ID: %s`
)

// newsletter
const (
	ErrNewsletterEmailAlreadySubscribed = `email already subscribed: %s`
	ErrNewsletterEmailNotYetSubscribed  = `email not yet subscribed: %s`
	ErrNewsletterInvalidEmail           = `email is invalid: %s`
	ErrNewsletterPublisherNotFound      = "one of the publishers sent were not found"
)
