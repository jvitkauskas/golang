package main

import "github.com/shopspring/decimal"

var data = map[string]decimal.Decimal{
	"Vilnius":  decimal.NewFromFloat(8.1),
	"Kaunas":   decimal.NewFromFloat(11.5),
	"Klaipėda": decimal.NewFromFloat(12.75),
}

func CityExists(city string) bool {
	_, ok := data[city]
	return ok
}

func GetTemperatureForCity(city string) (decimal.Decimal, bool) {
	temperature, ok := data[city]
	return temperature, ok
}

func SetTemperatureForCity(city string, temperature decimal.Decimal) {
	data[city] = temperature
}
