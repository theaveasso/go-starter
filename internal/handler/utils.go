package handler

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func newApiError(status int, err error) APIError {
	return APIError{
		StatusCode: status, Message: err.Error(),
	}
}

func (e APIError) Error() string {
	return e.Message
}

type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, v Envelope, headers http.Header) error {
	for key, value := range headers {
		w.Header()[key] = value
	}

  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
