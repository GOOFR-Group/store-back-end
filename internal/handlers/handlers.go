package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GOOFR-Group/store-back-end/internal/logging"
)

const (
	contentTypeJSON   = "application/json"
	contentTypeHeader = "Content-Type"
)

type header [2]string

func (h header) Key() string {
	return h[0]
}
func (h header) Value() string {
	return h[1]
}

func writeResponse(w http.ResponseWriter, response interface{}, statusCode int, headers ...header) error {
	message, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("failed to encode message as JSON response: %s", err.Error())
	}
	for _, h := range headers {
		w.Header().Set(h.Key(), h.Value())
	}
	w.WriteHeader(statusCode)
	if _, err = w.Write(message); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("failed to write message to the response body: %s", err.Error())
	}

	return nil
}

func logInternalError(handlerName string, errorName string, err error) {
	logging.AppLogger.Error().Str("handler", handlerName).Str("error", ErrWritingResponse).Msg(err.Error())
}
