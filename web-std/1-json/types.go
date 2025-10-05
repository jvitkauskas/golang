package main

import "github.com/shopspring/decimal"

type ErrorResponse struct {
	Message string `json:"message"`
}

type WeatherResponse struct {
	City        string          `json:"city"`
	Temperature decimal.Decimal `json:"temperature"`
}

type WeatherUpdateRequest struct {
	Temperature decimal.NullDecimal `json:"temperature"`
}
