package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerSearch = "search"

// GetSearchGame handles the /searchGame Get endpoint
func (*StoreImpl) GetSearchGame(w http.ResponseWriter, r *http.Request, params oapi.GetSearchGameParams) {
}

// GetSearchTag handles the /searchTag Get endpoint
func (*StoreImpl) GetSearchTag(w http.ResponseWriter, r *http.Request, params oapi.GetSearchTagParams) {
}

// PostSearchHistory handles the /searchHistory Post endpoint
func (*StoreImpl) PostSearchHistory(w http.ResponseWriter, r *http.Request, params oapi.PostSearchHistoryParams) {
}

// GetSearchHistory handles the /searchHistory Get endpoint
func (*StoreImpl) GetSearchHistory(w http.ResponseWriter, r *http.Request, params oapi.GetSearchHistoryParams) {
}
