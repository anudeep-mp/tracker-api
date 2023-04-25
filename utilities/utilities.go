package utilities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anudeep-mp/tracker/model"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ResponseWrapper(w http.ResponseWriter, status int, success bool, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	response := model.APIResponse{
		Success: success,
		Message: message,
		Result:  result,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
