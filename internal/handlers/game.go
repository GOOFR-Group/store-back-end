package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/core"
	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerGame = "game"

// PostGame handles the /game Post endpoint
func (*StoreImpl) PostGame(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerGame, err)
		return
	}

	req := oapi.PostGameJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerGame, err)
		return
	}

	response, err := core.PostGame(req)
	switch err {
	case nil:
	case core.ErrPublisherNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrPublisherNotFound, req.IdPublisher))
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerGame, err)
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// GetGame handles the /game Get endpoint
func (*StoreImpl) GetGame(w http.ResponseWriter, r *http.Request, params oapi.GetGameParams) {
	response, err := core.GetGame(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, *params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// PutGame handles the /game Put endpoint
func (*StoreImpl) PutGame(w http.ResponseWriter, r *http.Request, params oapi.PutGameParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerGame, err)
		return
	}

	req := oapi.PutGameJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerGame, err)
		return
	}

	response, err := core.PutGame(params, req)
	switch err {
	case nil:
	case core.ErrPublisherNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrPublisherNotFound, req.IdPublisher))
		return
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// DeleteGame handles the /game Delete endpoint
func (*StoreImpl) DeleteGame(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameParams) {
	response, err := core.DeleteGame(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// PostGameTag handles the /gameTag Post endpoint
func (*StoreImpl) PostGameTag(w http.ResponseWriter, r *http.Request, params oapi.PostGameTagParams) {
	err := core.PostGameTag(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	case core.ErrTagNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrTagNotFound, params.TagID))
		return
	case core.ErrObjectAlreadyCreated:
		writeConflict(w, handlerGame, ErrGameAlreadyContainsTag)
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeCreated(w)
}

// GetGameTag handles the /gameTag Get endpoint
func (*StoreImpl) GetGameTag(w http.ResponseWriter, r *http.Request, params oapi.GetGameTagParams) {
	response, err := core.GetGameTag(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// DeleteGameTag handles the /gameTag Delete endpoint
func (*StoreImpl) DeleteGameTag(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameTagParams) {
	response, err := core.DeleteGameTag(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	case core.ErrTagNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrTagNotFound, params.TagID))
		return
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerGame, ErrGameNotYetContainTag)
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// PostGameImage handles the /gameImage Post endpoint
func (*StoreImpl) PostGameImage(w http.ResponseWriter, r *http.Request, params oapi.PostGameImageParams) {
	err := core.PostGameImage(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.GameID))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeCreated(w)
}

// GetGameImage handles the /gameImage Get endpoint
func (*StoreImpl) GetGameImage(w http.ResponseWriter, r *http.Request, params oapi.GetGameImageParams) {
	response, err := core.GetGameImage(params)
	switch err {
	case nil:
	case core.ErrGameNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrGameNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}

// DeleteGameImage handles the /gameImage Delete endpoint
func (*StoreImpl) DeleteGameImage(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameImageParams) {
	response, err := core.DeleteGameImage(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerGame, fmt.Sprintf(ErrImageNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerGame, err)
		return
	}

	writeOK(w, handlerGame, response)
}
