package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerPublisher = "publisher"

// PostPublisher handles the /publisher Post endpoint
func (*StoreImpl) PostPublisher(w http.ResponseWriter, r *http.Request) {
}

// GetPublisher handles the /publisher Get endpoint
func (*StoreImpl) GetPublisher(w http.ResponseWriter, r *http.Request, params oapi.GetPublisherParams) {
}

// PutPublisher handles the /publisher Put endpoint
func (*StoreImpl) PutPublisher(w http.ResponseWriter, r *http.Request, params oapi.PutPublisherParams) {
}

// DeletePublisher handles the /publisher Delete endpoint
func (*StoreImpl) DeletePublisher(w http.ResponseWriter, r *http.Request, params oapi.DeletePublisherParams) {
}

// GetPublisherGames handles the /publisher Get endpoint
func (*StoreImpl) GetPublisherGames(w http.ResponseWriter, r *http.Request, params oapi.GetPublisherGamesParams) {
}
