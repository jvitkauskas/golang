package main

import (
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	repo := NewRepository()
	handler := NewWeatherHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather/{city}", handler.GetCityTemperature)
	mux.HandleFunc("PUT /weather/{city}", handler.PutCityTemperature)
	mux.HandleFunc("POST /weather/{city}", handler.PostCityTemperature)

	port := ":8080"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, mux)
}
