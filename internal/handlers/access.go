package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerAccess = "access"

// PostLogin handles the /login Post endpoint
func (*StoreImpl) PostLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	req := oapi.PostLoginJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	response, err := core.PostLogin(req)
	switch err {
	case nil:
	case core.ErrClientInactive:
		writeConflict(w, handlerAccess, ErrAccessinactive)
		return
	case core.ErrClientNotFound, core.ErrPasswordRequired, core.ErrIncorrectPassword, core.ErrObjectNotFound:
		writeConflict(w, handlerAccess, ErrAccessIncorrect)
		return
	default:
		writeInternalServerError(w, handlerAccess, err)
		return
	}

	writeOK(w, handlerAccess, response)
}

// PostRegister handles the /register Post endpoint
func (*StoreImpl) PostRegister(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	req := oapi.PostRegisterJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	response, err := core.PostRegister(req)
	switch err {
	case nil:
	case core.ErrInvalidEmail:
		writeConflict(w, handlerAccess, ErrAccessInvalidEmail)
		return
	case core.ErrPasswordRequired:
		writeConflict(w, handlerAccess, ErrAccessPasswordRequired)
		return
	case core.ErrInvalidPassword:
		writeConflict(w, handlerAccess, ErrAccessInvalidPassword)
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerAccess, ErrAccessAlreadyCreated)
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerAccess, err)
		return
	default:
		writeInternalServerError(w, handlerAccess, err)
		return
	}

	writeOK(w, handlerAccess, response)
}

// PutAccess handles the /access Put endpoint
func (*StoreImpl) PutAccess(w http.ResponseWriter, r *http.Request, params oapi.PutAccessParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	req := oapi.PutAccessJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerAccess, err)
		return
	}

	err = core.PutAccess(params, req)
	switch err {
	case nil:
	case core.ErrInvalidEmail:
		writeConflict(w, handlerAccess, ErrAccessInvalidEmail)
		return
	case core.ErrPasswordRequired:
		writeConflict(w, handlerAccess, ErrAccessPasswordRequired)
		return
	case core.ErrInvalidPassword:
		writeConflict(w, handlerAccess, ErrAccessInvalidPassword)
		return
	case core.ErrClientNotFound:
		writeConflict(w, handlerAccess, fmt.Sprintf(ErrClientNotFound, params.ClientID))
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerAccess, ErrAccessAlreadyCreated)
		return
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerAccess, err)
		return
	default:
		writeInternalServerError(w, handlerAccess, err)
		return
	}

	writeNoContent(w)
}
