package handlers

import (
	"net/http"

	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerCart = "cart"

// PostCart handles the /cart Post endpoint
func (*StoreImpl) PostCart(w http.ResponseWriter, r *http.Request, params oapi.PostCartParams) {
}

// GetCart handles the /cart Get endpoint
func (*StoreImpl) GetCart(w http.ResponseWriter, r *http.Request, params oapi.GetCartParams) {
}

// DeleteCart handles the /cart Delete endpoint
func (*StoreImpl) DeleteCart(w http.ResponseWriter, r *http.Request, params oapi.DeleteCartParams) {
}

// GetCartPurchase handles the /cartPurchase Get endpoint
func (*StoreImpl) GetCartPurchase(w http.ResponseWriter, r *http.Request, params oapi.GetCartPurchaseParams) {
}
