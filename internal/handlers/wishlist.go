package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerWishlist = "wishlist"

// PostWishlist handles the /wishlist Post endpoint
func (*StoreImpl) PostWishlist(w http.ResponseWriter, r *http.Request, params oapi.PostWishlistParams) {
}

// GetWishlist handles the /wishlist Get endpoint
func (*StoreImpl) GetWishlist(w http.ResponseWriter, r *http.Request, params oapi.GetWishlistParams) {
}

// DeleteWishlist handles the /wishlist Delete endpoint
func (*StoreImpl) DeleteWishlist(w http.ResponseWriter, r *http.Request, params oapi.DeleteWishlistParams) {
}
