package handlers

import (
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/logging"
)

const handlerHealth = "health"

func (*StoreImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	if err := writeResponse(w, "I'm fine, thanks for asking!", http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logging.AppLogger.Error().Msgf(ErrWritingResponse, handlerHealth, err.Error())
	}
}
