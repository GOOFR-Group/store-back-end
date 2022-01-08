package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goofr-group/store-back-end/internal/logging"
	"github.com/goofr-group/store-back-end/internal/oapi"
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

// writeOK 200
func writeOK(w http.ResponseWriter, handlerName string, response interface{}) {
	if err := writeResponse(w, response, http.StatusOK, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerName, ErrWritingResponse, err)
	}
}

// writeCreated 201
func writeCreated(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

// writeNoContent 204
func writeNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// writeBadRequest 400
func writeBadRequest(w http.ResponseWriter, handlerName string, err error) {
	response := oapi.ErrorSchema{Error: fmt.Sprintf(ErrParsingRequest, err.Error())}
	if err = writeResponse(w, response, http.StatusBadRequest, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerName, ErrWritingResponse, err)
	}
}

// writeNotFound 404
func writeNotFound(w http.ResponseWriter, handlerName string, errorMessage string) {
	response := oapi.ErrorSchema{Error: errorMessage}
	if err := writeResponse(w, response, http.StatusNotFound, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerName, ErrWritingResponse, err)
	}
}

// writeConflict 409
func writeConflict(w http.ResponseWriter, handlerName string, errorMessage string) {
	response := oapi.ErrorSchema{Error: errorMessage}
	if err := writeResponse(w, response, http.StatusConflict, header{contentTypeHeader, contentTypeJSON}); err != nil {
		logInternalError(handlerName, ErrWritingResponse, err)
	}
}

// writeInternalServerError 500
func writeInternalServerError(w http.ResponseWriter, handlerName string, err error) {
	logInternalError(handlerName, ErrInternalServer, err)
	w.WriteHeader(http.StatusInternalServerError)
}
