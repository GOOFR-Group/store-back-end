package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
)

const handlerStatistics = "statistics"

// GetTopReviews handles the /topReviews Get endpoint
func (*StoreImpl) GetTopReviews(w http.ResponseWriter, r *http.Request) {
	response, err := core.GetTopReviews()
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerStatistics, err)
		return
	}

	writeOK(w, handlerStatistics, response)
}
