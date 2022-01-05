package handlers

import (
	"net/http"
)

const handlerHealth = "health"

// GetHealth handles the healt check endpoint
func (*StoreImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	if err := writeResponse(w, "I'm fine, thanks for asking!", http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerHealth, ErrInternalServer, err)
	}
}
