package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJSON[T any](r *http.Request) (T, error) {
	var v T

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&v); err != nil {
		return v, err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return v, errors.New("body must only contain a single JSON value")
	}

	return v, nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}