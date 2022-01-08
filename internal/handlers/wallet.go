package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerWallet = "wallet"

// GetWallet handles the /wallet Get endpoint
func (*StoreImpl) GetWallet(w http.ResponseWriter, r *http.Request, params oapi.GetWalletParams) {
	response, err := core.GetWallet(params)
	switch err {
	case nil:
	case core.ErrClientNotFound:
		writeNotFound(w, handlerWallet, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerWallet, err)
		return
	default:
		writeInternalServerError(w, handlerWallet, err)
		return
	}

	writeOK(w, handlerWallet, response)
}

// PutWallet handles the /wallet Put endpoint
func (*StoreImpl) PutWallet(w http.ResponseWriter, r *http.Request, params oapi.PutWalletParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerWallet, err)
		return
	}

	req := oapi.PutWalletJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerWallet, err)
		return
	}

	response, err := core.PutWallet(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerWallet, fmt.Sprintf(ErrWalletNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerWallet, err)
		return
	}

	writeOK(w, handlerWallet, response)
}

// GetAddBalance handles the /addBalance Get endpoint
func (*StoreImpl) GetAddBalance(w http.ResponseWriter, r *http.Request, params oapi.GetAddBalanceParams) {
	response, err := core.GetAddBalance(params)
	switch err {
	case nil:
	case core.ErrInvalidAmount:
		writeConflict(w, handlerWallet, ErrWalletInvalidAmount)
		return
	case core.ErrClientNotFound:
		writeNotFound(w, handlerWallet, fmt.Sprintf(ErrClientNotFound, params.Id))
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerWallet, err)
		return
	default:
		writeInternalServerError(w, handlerWallet, err)
		return
	}

	writeOK(w, handlerWallet, response)
}
