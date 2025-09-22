package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(APIError{
		Code:    code,
		Message: message,
	})
	if err != nil {
		log.Printf("Failed to write error response: %v", err)
	}
}
