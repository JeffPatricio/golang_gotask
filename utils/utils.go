package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	ToJson(w, struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
	}{
		Message: err,
		Success: false,
	})
}

func ToJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func IsEmpty(param string) bool {
	if param == "" {
		return true
	}
	return false
}
