package handlers

// general
const (
	// internal
	ErrWritingResponse = "failed to write response"
	ErrInternalServer  = "internal server error"

	// response
	ErrParsingRequest = "error parsing request: %s"
)

// game
const (
	ErrGameNotFound      = "game not found with ID: %s"
	ErrGameAlreadyBought = "client already has the given game in his library"
)

// game tag
const (
	ErrGameAlreadyContainsTag = "game already contains the given tag"
	ErrGameNotYetContainTag   = "game does not yet contain the given tag"
)

// game image
const (
	ErrImageNotFound = "image not found with ID: %s"
)

// publisher
const (
	ErrPublisherNotFound = "publisher not found with ID: %s"
)

// client
const (
	ErrClientNotFound = "client not found with ID: %s"
)

// access
const (
	ErrAccessinactive         = "inactive client"
	ErrAccessIncorrect        = "incorrect email or password"
	ErrAccessInvalidEmail     = "invalid email"
	ErrAccessPasswordRequired = "password not submitted"
	ErrAccessInvalidPassword  = "password must be at least 6 characters long, composed of letters with at least one digit and one special character"
	ErrAccessAlreadyCreated   = "email already registered"
)

// address
const (
	ErrAddressNotFound = "address not found with ID: %s"
)

// wallet
const (
	ErrWalletNotFound      = "wallet not found with ID: %s"
	ErrWalletInvalidAmount = "invalid amount"
)

// cart
const (
	ErrCartGameAlreadyAdded    = "client already has the given game in his cart"
	ErrCartGameNotAdded        = "client does not have this game in his cart"
	ErrCartInsufficientBalance = "client does not have enough balance"
	ErrCartEmpty               = "client has no games in his cart"
)

// wishlist
const (
	ErrWishlistGameAlreadyAdded = "client already has the given game in his wishlist"
	ErrWishlistGameNotAdded     = "client does not have this game in his wishlist"
)

// review
const (
	ErrReviewAlreadyCreated = "client has already reviewed the game"
	ErrReviewNotFound       = "review not found with ID: %s"
)

// tag
const (
	ErrTagNotFound = "tag not found with ID: %s"
)

// newsletter
const (
	ErrNewsletterEmailAlreadySubscribed = "email already subscribed: %s"
	ErrNewsletterEmailNotYetSubscribed  = "email not yet subscribed: %s"
	ErrNewsletterInvalidEmail           = "email is invalid: %s"
	ErrNewsletterPublisherNotFound      = "one of the publishers sent were not found"
)
