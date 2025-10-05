package main

import (
	"net/http"
)

type WeatherHandler struct {
	repo *Repository
}

func NewWeatherHandler(repo *Repository) *WeatherHandler {
	return &WeatherHandler{repo: repo}
}

func (s *WeatherHandler) GetCityTemperature(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")
	temperature, ok := s.repo.GetTemperatureForCity(city)
	if !ok {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "No such city"})
		return
	}

	WriteJSON(w, http.StatusOK, WeatherResponse{City: city, Temperature: temperature})
}

func (s *WeatherHandler) PutCityTemperature(w http.ResponseWriter, r *http.Request) {
	s.saveCityTemperature(w, r, true)
}

func (s *WeatherHandler) PostCityTemperature(w http.ResponseWriter, r *http.Request) {
	s.saveCityTemperature(w, r, false)
}

func (s *WeatherHandler) saveCityTemperature(w http.ResponseWriter, r *http.Request, checkExisting bool) {
	city := r.PathValue("city")

	if checkExisting && !s.repo.CityExists(city) {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "No such city"})
		return
	}

	request, err := ReadJSON[WeatherUpdateRequest](r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if !request.Temperature.Valid {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Message: "Invalid or no temperature provided"})
		return
	}

	s.repo.SetTemperatureForCity(city, request.Temperature.Decimal)

	WriteJSON(w, http.StatusOK, WeatherResponse{City: city, Temperature: request.Temperature.Decimal})
}
