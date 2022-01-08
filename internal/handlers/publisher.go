package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/core"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const handlerPublisher = "publisher"

// PostPublisher handles the /publisher Post endpoint
func (*StoreImpl) PostPublisher(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerPublisher, err)
		return
	}

	req := oapi.PostPublisherJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerPublisher, err)
		return
	}

	response, err := core.PostPublisher(req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeInternalServerError(w, handlerPublisher, err)
	default:
		writeInternalServerError(w, handlerPublisher, err)
		return
	}

	writeOK(w, handlerPublisher, response)
}

// GetPublisher handles the /publisher Get endpoint
func (*StoreImpl) GetPublisher(w http.ResponseWriter, r *http.Request, params oapi.GetPublisherParams) {
	response, err := core.GetPublisher(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerPublisher, fmt.Sprintf(ErrPublisherNotFound, *params.Id))
		return
	default:
		writeInternalServerError(w, handlerPublisher, err)
		return
	}

	writeOK(w, handlerPublisher, response)
}

// PutPublisher handles the /publisher Put endpoint
func (*StoreImpl) PutPublisher(w http.ResponseWriter, r *http.Request, params oapi.PutPublisherParams) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeBadRequest(w, handlerPublisher, err)
		return
	}

	req := oapi.PutPublisherJSONRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeBadRequest(w, handlerPublisher, err)
		return
	}

	response, err := core.PutPublisher(params, req)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerPublisher, fmt.Sprintf(ErrPublisherNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerPublisher, err)
		return
	}

	writeOK(w, handlerPublisher, response)
}

// DeletePublisher handles the /publisher Delete endpoint
func (*StoreImpl) DeletePublisher(w http.ResponseWriter, r *http.Request, params oapi.DeletePublisherParams) {
	response, err := core.DeletePublisher(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerPublisher, fmt.Sprintf(ErrPublisherNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerPublisher, err)
		return
	}

	writeOK(w, handlerPublisher, response)
}

// GetPublisherGames handles the /publisher Get endpoint
func (*StoreImpl) GetPublisherGames(w http.ResponseWriter, r *http.Request, params oapi.GetPublisherGamesParams) {
	response, err := core.GetPublisherGames(params)
	switch err {
	case nil:
	case core.ErrObjectNotFound:
		writeNotFound(w, handlerPublisher, fmt.Sprintf(ErrPublisherNotFound, params.Id))
		return
	default:
		writeInternalServerError(w, handlerPublisher, err)
		return
	}

	writeOK(w, handlerPublisher, response)
}
