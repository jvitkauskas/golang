package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	port := "8080"
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, mux)
}
