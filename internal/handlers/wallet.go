package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerWallet = "wallet"

// GetWallet handles the /wallet Get endpoint
func (*StoreImpl) GetWallet(w http.ResponseWriter, r *http.Request, params oapi.GetWalletParams) {
}

// PutWallet handles the /wallet Put endpoint
func (*StoreImpl) PutWallet(w http.ResponseWriter, r *http.Request, params oapi.PutWalletParams) {
}

// GetAddBalance handles the /addBalance Get endpoint
func (*StoreImpl) GetAddBalance(w http.ResponseWriter, r *http.Request, params oapi.GetAddBalanceParams) {
}
