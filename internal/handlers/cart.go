package handlers

import (
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerCart = "cart"

// PostCart handles the /cart Post endpoint
func (*StoreImpl) PostCart(w http.ResponseWriter, r *http.Request, params oapi.PostCartParams) {
	err := core.PostCart(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrGameNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	case core.ErrGameAlreadyBought:
		writeConflict(w, handlerCart, ErrGameAlreadyBought)
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerCart, ErrCartGameAlreadyAdded)
		return
	default:
		writeInternalServerError(w, handlerCart, err)
		return
	}

	writeCreated(w)
}

// GetCart handles the /cart Get endpoint
func (*StoreImpl) GetCart(w http.ResponseWriter, r *http.Request, params oapi.GetCartParams) {
	response, err := core.GetCart(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerCart, err)
		return
	}

	writeOK(w, handlerCart, response)
}

// DeleteCart handles the /cart Delete endpoint
func (*StoreImpl) DeleteCart(w http.ResponseWriter, r *http.Request, params oapi.DeleteCartParams) {
	response, err := core.DeleteCart(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrGameNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrGameNotFound, *params.GameID))
		return
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerCart, ErrCartGameNotAdded)
		return
	default:
		writeInternalServerError(w, handlerCart, err)
		return
	}

	writeOK(w, handlerCart, response)
}

// GetCartPurchase handles the /cartPurchase Get endpoint
func (*StoreImpl) GetCartPurchase(w http.ResponseWriter, r *http.Request, params oapi.GetCartPurchaseParams) {
	response, err := core.GetCartPurchase(params)
	switch err {
	case nil:
	case core.ErrInvalidAmount:
		writeConflict(w, handlerCart, ErrCartInsufficientBalance)
		return
	case core.ErrClientNotFound:
		writeNotFound(w, handlerCart, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerCart, ErrCartEmpty)
		return
	case core.ErrInvoiceHeaderNotFound:
		writeInternalServerError(w, handlerCart, err)
		return
	default:
		writeInternalServerError(w, handlerCart, err)
		return
	}

	writeOK(w, handlerCart, response)
}
