package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "Application/json")
	return json.NewEncoder(w).Encode(v)
}
