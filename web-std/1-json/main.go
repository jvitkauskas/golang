package main

import (
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather/{city}", GetCityTemperature)
	mux.HandleFunc("PUT /weather/{city}", PutCityTemperature)
	mux.HandleFunc("POST /weather/{city}", PostCityTemperature)

	port := ":8080"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, mux)
}
