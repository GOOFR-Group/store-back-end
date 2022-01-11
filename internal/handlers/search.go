package handlers

import (
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerSearch = "search"

// GetSearchGame handles the /searchGame Get endpoint
func (*StoreImpl) GetSearchGame(w http.ResponseWriter, r *http.Request, params oapi.GetSearchGameParams) {
	response, err := core.GetSearchGame(params)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerSearch, err)
		return
	}

	writeOK(w, handlerSearch, response)
}

// GetSearchTag handles the /searchTag Get endpoint
func (*StoreImpl) GetSearchTag(w http.ResponseWriter, r *http.Request, params oapi.GetSearchTagParams) {
	response, err := core.GetSearchTag(params)
	switch err {
	case nil:
	default:
		writeInternalServerError(w, handlerSearch, err)
		return
	}

	writeOK(w, handlerSearch, response)
}

// PostSearchHistory handles the /searchHistory Post endpoint
func (*StoreImpl) PostSearchHistory(w http.ResponseWriter, r *http.Request, params oapi.PostSearchHistoryParams) {
	err := core.PostSearchHistory(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerSearch, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrGameNotFound:
		writeNotFound(w, handlerSearch, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	default:
		writeInternalServerError(w, handlerSearch, err)
		return
	}

	writeCreated(w)
}

// GetSearchHistory handles the /searchHistory Get endpoint
func (*StoreImpl) GetSearchHistory(w http.ResponseWriter, r *http.Request, params oapi.GetSearchHistoryParams) {
	response, err := core.GetSearchHistory(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerSearch, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerSearch, err)
		return
	}

	writeOK(w, handlerSearch, response)
}
