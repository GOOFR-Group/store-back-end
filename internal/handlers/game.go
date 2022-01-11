package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerGame = "game"

// PostGame handles the /game Post endpoint
func (*StoreImpl) PostGame(w http.ResponseWriter, r *http.Request) {
}

// GetGame handles the /game Get endpoint
func (*StoreImpl) GetGame(w http.ResponseWriter, r *http.Request, params oapi.GetGameParams) {
}

// PutGame handles the /game Put endpoint
func (*StoreImpl) PutGame(w http.ResponseWriter, r *http.Request, params oapi.PutGameParams) {
}

// DeleteGame handles the /game Delete endpoint
func (*StoreImpl) DeleteGame(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameParams) {
}

// PostGameTag handles the /gameTag Post endpoint
func (*StoreImpl) PostGameTag(w http.ResponseWriter, r *http.Request, params oapi.PostGameTagParams) {
}

// GetGameTag handles the /gameTag Get endpoint
func (*StoreImpl) GetGameTag(w http.ResponseWriter, r *http.Request, params oapi.GetGameTagParams) {
}

// DeleteGameTag handles the /gameTag Delete endpoint
func (*StoreImpl) DeleteGameTag(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameTagParams) {
}

// PostGameImage handles the /gameImage Post endpoint
func (*StoreImpl) PostGameImage(w http.ResponseWriter, r *http.Request, params oapi.PostGameImageParams) {
}

// GetGameImage handles the /gameImage Get endpoint
func (*StoreImpl) GetGameImage(w http.ResponseWriter, r *http.Request, params oapi.GetGameImageParams) {
}

// DeleteGameImage handles the /gameImage Delete endpoint
func (*StoreImpl) DeleteGameImage(w http.ResponseWriter, r *http.Request, params oapi.DeleteGameImageParams) {
}
