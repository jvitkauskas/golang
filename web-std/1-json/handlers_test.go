package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGetCityTemperature(t *testing.T) {
	repo := newTestRepository()
	handler := NewWeatherHandler(repo)

	t.Run("existing city", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/weather/Vilnius", nil)
		req.SetPathValue("city", "Vilnius")
		rr := httptest.NewRecorder()

		handler.GetCityTemperature(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}

		var body WeatherResponse
		if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
			t.Errorf("unexpected error decoding body: %v", err)
		}

		if body.City != "Vilnius" {
			t.Errorf("expected city %v, got %v", "Vilnius", body.City)
		}
	})

	t.Run("non-existing city", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/weather/London", nil)
		req.SetPathValue("city", "London")
		rr := httptest.NewRecorder()

		handler.GetCityTemperature(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})
}

func TestPutCityTemperature(t *testing.T) {
	repo := newTestRepository()
	handler := NewWeatherHandler(repo)

	t.Run("existing city", func(t *testing.T) {
		body := `{"temperature": "15.5"}`
		req := httptest.NewRequest(http.MethodPut, "/weather/Vilnius", bytes.NewBufferString(body))
		req.SetPathValue("city", "Vilnius")
		rr := httptest.NewRecorder()

		handler.PutCityTemperature(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}

		var resBody WeatherResponse
		if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
			t.Errorf("unexpected error decoding body: %v", err)
		}

		if resBody.City != "Vilnius" {
			t.Errorf("expected city %v, got %v", "Vilnius", resBody.City)
		}

		expectedTemp := decimal.NewFromFloat(15.5)
		if !resBody.Temperature.Equal(expectedTemp) {
			t.Errorf("expected temperature %v, got %v", expectedTemp, resBody.Temperature)
		}
	})

	t.Run("non-existing city", func(t *testing.T) {
		body := `{"temperature": "15.5"}`
		req := httptest.NewRequest(http.MethodPut, "/weather/London", bytes.NewBufferString(body))
		req.SetPathValue("city", "London")
		rr := httptest.NewRecorder()

		handler.PutCityTemperature(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		body := `{"temperature": "invalid"}`
		req := httptest.NewRequest(http.MethodPut, "/weather/Vilnius", bytes.NewBufferString(body))
		req.SetPathValue("city", "Vilnius")
		rr := httptest.NewRecorder()

		handler.PutCityTemperature(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})
}

func TestPostCityTemperature(t *testing.T) {
	repo := newTestRepository()
	server := NewWeatherHandler(repo)

	t.Run("new city", func(t *testing.T) {
		body := `{"temperature": "15.5"}`
		req := httptest.NewRequest(http.MethodPost, "/weather/London", bytes.NewBufferString(body))
		req.SetPathValue("city", "London")
		rr := httptest.NewRecorder()

		server.PostCityTemperature(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}

		var resBody WeatherResponse
		if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
			t.Errorf("unexpected error decoding body: %v", err)
		}

		if resBody.City != "London" {
			t.Errorf("expected city %v, got %v", "London", resBody.City)
		}

		expectedTemp := decimal.NewFromFloat(15.5)
		if !resBody.Temperature.Equal(expectedTemp) {
			t.Errorf("expected temperature %v, got %v", expectedTemp, resBody.Temperature)
		}
	})

	t.Run("existing city", func(t *testing.T) {
		body := `{"temperature": "15.5"}`
		req := httptest.NewRequest(http.MethodPost, "/weather/Vilnius", bytes.NewBufferString(body))
		req.SetPathValue("city", "Vilnius")
		rr := httptest.NewRecorder()

		server.PostCityTemperature(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		body := `{"temperature": "invalid"}`
		req := httptest.NewRequest(http.MethodPost, "/weather/London", bytes.NewBufferString(body))
		req.SetPathValue("city", "London")
		rr := httptest.NewRecorder()

		server.PostCityTemperature(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})
}
