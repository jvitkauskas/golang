package main

import (
	"encoding/json"
	"net/http"
)

func ReadJSON[T any](r *http.Request, v *T) *T {
	json.NewDecoder(r.Body).Decode(v)
	return v
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
