package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type APIResponse struct {
	Success bool        `json:"isSuccess"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseWrapper(w http.ResponseWriter, status int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	response := APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
