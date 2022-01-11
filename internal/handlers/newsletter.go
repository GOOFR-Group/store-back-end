package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerNewsletter = "newsletter"

// PostNewsletter handles the /newsletter Post endpoint
func (*StoreImpl) PostNewsletter(w http.ResponseWriter, r *http.Request, params oapi.PostNewsletterParams) {
	err := core.PostNewsletter(params)
	switch err {
	case nil:
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerNewsletter, fmt.Sprintf(ErrNewsletterEmailAlreadySubscribed, params.Email))
		return
	case core.ErrInvalidEmail:
		writeConflict(w, handlerNewsletter, fmt.Sprintf(ErrNewsletterInvalidEmail, params.Email))
		return
	default:
		writeInternalServerError(w, handlerNewsletter, err)
		return
	}

	writeCreated(w)
}

// GetNewsletter handles the /newsletter Get endpoint
func (*StoreImpl) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	response, err := core.GetNewsletter()
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerNewsletter, err)
		return
	}

	writeOK(w, handlerNewsletter, response)
}

// DeleteNewsletter handles the /newsletter Delete endpoint
func (*StoreImpl) DeleteNewsletter(w http.ResponseWriter, r *http.Request, params oapi.DeleteNewsletterParams) {
	response, err := core.DeleteNewsletter(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerNewsletter, fmt.Sprintf(ErrNewsletterEmailNotYetSubscribed, params.Email))
		return
	default:
		writeInternalServerError(w, handlerNewsletter, err)
		return
	}

	writeOK(w, handlerNewsletter, response)
}

// PostSendNewsletter handles the /sendNewsletter Post endpoint
func (*StoreImpl) PostSendNewsletter(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerNewsletter, err)
		return
	}

	req := oapi.PostSendNewsletterJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerNewsletter, err)
		return
	}

	err = core.PostSendNewsletter(req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerNewsletter, ErrNewsletterPublisherNotFound)
		return
	default:
		writeInternalServerError(w, handlerNewsletter, err)
		return
	}

	writeCreated(w)
}
