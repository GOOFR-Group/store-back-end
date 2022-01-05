package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerNewsletter = "newsletter"

// PostNewsletter handles the /newsletter Post endpoint
func (*StoreImpl) PostNewsletter(w http.ResponseWriter, r *http.Request, params oapi.PostNewsletterParams) {
}

// GetNewsletter handles the /newsletter Get endpoint
func (*StoreImpl) GetNewsletter(w http.ResponseWriter, r *http.Request) {
}

// DeleteNewsletter handles the /newsletter Delete endpoint
func (*StoreImpl) DeleteNewsletter(w http.ResponseWriter, r *http.Request, params oapi.DeleteNewsletterParams) {
}

// PostSendNewsletter handles the /sendNewsletter Post endpoint
func (*StoreImpl) PostSendNewsletter(w http.ResponseWriter, r *http.Request) {
}
