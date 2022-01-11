package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerReview = "review"

// PostReview handles the /review Post endpoint
func (*StoreImpl) PostReview(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerReview, err)
		return
	}

	req := oapi.PostReviewJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerReview, err)
		return
	}

	response, err := core.PostReview(req)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerReview, fmt.Sprintf(ErrGameNotFound, req.IdGame))
		return
	case core.ErrClientNotFound:
		writeNotFound(w, handlerReview, fmt.Sprintf(ErrClientNotFound, req.IdClient))
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerReview, ErrReviewAlreadyCreated)
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerReview, err)
		return
	default:
		writeInternalServerError(w, handlerReview, err)
		return
	}

	writeOK(w, handlerReview, response)
}

// PutReview handles the /review Put endpoint
func (*StoreImpl) PutReview(w http.ResponseWriter, r *http.Request, params oapi.PutReviewParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerReview, err)
		return
	}

	req := oapi.PutReviewJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerReview, err)
		return
	}

	response, err := core.PutReview(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerReview, fmt.Sprintf(ErrReviewNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerReview, err)
		return
	}

	writeOK(w, handlerReview, response)
}

// DeleteReview handles the /review Delete endpoint
func (*StoreImpl) DeleteReview(w http.ResponseWriter, r *http.Request, params oapi.DeleteReviewParams) {
	response, err := core.DeleteReview(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerReview, fmt.Sprintf(ErrReviewNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerReview, err)
		return
	}

	writeOK(w, handlerReview, response)
}

// GetGameReviews handles the /gameReviews Get endpoint
func (*StoreImpl) GetGameReviews(w http.ResponseWriter, r *http.Request, params oapi.GetGameReviewsParams) {
	response, err := core.GetGameReviews(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerReview, fmt.Sprintf(ErrGameNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerReview, err)
		return
	}

	writeOK(w, handlerReview, response)
}
