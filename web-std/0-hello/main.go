package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Handle root path with optional name query parameter
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello %s", name)
	})

	// Handle name parameter in the path
	mux.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "Hello %s", name)
	})

	port := "8080"
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, mux)
}
