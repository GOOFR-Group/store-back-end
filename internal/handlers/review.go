package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerReview = "review"

// PostReview handles the /review Post endpoint
func (*StoreImpl) PostReview(w http.ResponseWriter, r *http.Request) {
}

// PutReview handles the /review Put endpoint
func (*StoreImpl) PutReview(w http.ResponseWriter, r *http.Request, params oapi.PutReviewParams) {
}

// DeleteReview handles the /review Delete endpoint
func (*StoreImpl) DeleteReview(w http.ResponseWriter, r *http.Request, params oapi.DeleteReviewParams) {
}

// GetGameReviews handles the /gameReviews Get endpoint
func (*StoreImpl) GetGameReviews(w http.ResponseWriter, r *http.Request, params oapi.GetGameReviewsParams) {
}
