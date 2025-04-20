package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data any `json:"data"`
}

func RespondJSON(w http.ResponseWriter, status int, success bool, data any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(response{Data: data})
}
