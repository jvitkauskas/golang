package main

import (
	"net/http"
)

func GetCityTemperature(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")
	temperature, ok := GetTemperatureForCity(city)
	if !ok {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "No such city"})
		return
	}

	WriteJSON(w, http.StatusOK, WeatherResponse{City: city, Temperature: temperature})
}

func PutCityTemperature(w http.ResponseWriter, r *http.Request) {
	saveCityTemperature(w, r, true)
}

func PostCityTemperature(w http.ResponseWriter, r *http.Request) {
	saveCityTemperature(w, r, false)
}

func saveCityTemperature(w http.ResponseWriter, r *http.Request, checkExisting bool) {
	city := r.PathValue("city")

	if checkExisting && !CityExists(city) {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "No such city"})
		return
	}

	request := ReadJSON(r, &WeatherUpdateRequest{})
	if !request.Temperature.Valid {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "Invalid or no temperature provided"})
		return
	}

	SetTemperatureForCity(city, request.Temperature.Decimal)

	WriteJSON(w, http.StatusOK, WeatherResponse{City: city, Temperature: request.Temperature.Decimal})
}
