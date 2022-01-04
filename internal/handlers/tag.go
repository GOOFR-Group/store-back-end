package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

const handlerTag = "tag"

// PostTag handles the /tag Post endpoint
func (*StoreImpl) PostTag(w http.ResponseWriter, r *http.Request) {
	// logging.AppLogger.Error().Str(logHandlerKey, handlerTag).Str(logErrorKey, ErrWritingResponse).Msg(err.Error())
}

// GetTag handles the /tag Get endpoint
func (*StoreImpl) GetTag(w http.ResponseWriter, r *http.Request, params oapi.GetTagParams) {
}

// PutTag handles the /tag Put endpoint
func (*StoreImpl) PutTag(w http.ResponseWriter, r *http.Request, params oapi.PutTagParams) {
}

// DeleteTag handles the /tag Delete endpoint
func (*StoreImpl) DeleteTag(w http.ResponseWriter, r *http.Request, params oapi.DeleteTagParams) {
}
