package handlers

import (
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerWishlist = "wishlist"

// PostWishlist handles the /wishlist Post endpoint
func (*StoreImpl) PostWishlist(w http.ResponseWriter, r *http.Request, params oapi.PostWishlistParams) {
	err := core.PostWishlist(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerWishlist, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrGameNotFound:
		writeNotFound(w, handlerWishlist, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	case core.ErrGameAlreadyBought:
		writeConflict(w, handlerWishlist, ErrGameAlreadyBought)
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerWishlist, ErrWishlistGameAlreadyAdded)
		return
	default:
		writeInternalServerError(w, handlerWishlist, err)
		return
	}

	writeCreated(w)
}

// GetWishlist handles the /wishlist Get endpoint
func (*StoreImpl) GetWishlist(w http.ResponseWriter, r *http.Request, params oapi.GetWishlistParams) {
	response, err := core.GetWishlist(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerWishlist, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerWishlist, err)
		return
	}

	writeOK(w, handlerWishlist, response)
}

// DeleteWishlist handles the /wishlist Delete endpoint
func (*StoreImpl) DeleteWishlist(w http.ResponseWriter, r *http.Request, params oapi.DeleteWishlistParams) {
	response, err := core.DeleteWishlist(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerWishlist, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrGameNotFound:
		writeNotFound(w, handlerWishlist, fmt.Sprintf(ErrGameNotFound, *params.GameID))
		return
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerWishlist, ErrWishlistGameNotAdded)
		return
	default:
		writeInternalServerError(w, handlerWishlist, err)
		return
	}

	writeOK(w, handlerWishlist, response)
}
