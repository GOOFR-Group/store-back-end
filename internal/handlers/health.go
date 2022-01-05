package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/logging"
)

const handlerHealth = "health"

// GetHealth handles the healt check endpoint
func (*StoreImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	if err := writeResponse(w, "I'm fine, thanks for asking!", http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logging.AppLogger.Error().Str(logHandlerKey, handlerHealth).Str(logErrorKey, ErrWritingResponse).Msg(err.Error())
	}
}
