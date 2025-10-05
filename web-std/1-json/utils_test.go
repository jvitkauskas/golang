package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestReadJSON(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		// Create a request with a valid JSON body
		body := `{"temperature": "12.3"}`
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))

		// Call the function
		data, err := ReadJSON[WeatherUpdateRequest](req)

		// Check for errors
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Check the data
		expected := decimal.NewFromFloat(12.3)
		if !data.Temperature.Decimal.Equal(expected) {
			t.Errorf("expected %v, got %v", expected, data.Temperature.Decimal)
		}
	})

	t.Run("unknown field", func(t *testing.T) {
		// Create a request with an unknown field in the JSON body
		body := `{"temperature": "12.3", "unknown": "field"}`
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))

		// Call the function
		_, err := ReadJSON[WeatherUpdateRequest](req)

		// Check for errors
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("multiple json objects", func(t *testing.T) {
		// Create a request with multiple JSON objects in the body
		body := `{"temperature": "12.3"}{"temperature": "15.5"}`
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))

		// Call the function
		_, err := ReadJSON[WeatherUpdateRequest](req)

		// Check for errors
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		// Create a request with invalid JSON in the body
		body := `{"temperature": "12.3"`
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))

		// Call the function
		_, err := ReadJSON[WeatherUpdateRequest](req)

		// Check for errors
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}

func TestWriteJSON(t *testing.T) {
	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create some data to write
	data := WeatherResponse{
		City:        "Vilnius",
		Temperature: decimal.NewFromFloat(12.3),
	}

	// Call the function
	err := WriteJSON(rr, http.StatusOK, data)

	// Check for errors
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	if rr.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("expected content type %v, got %v", expectedContentType, rr.Header().Get("Content-Type"))
	}

	// Check the body
	var body WeatherResponse
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Errorf("unexpected error decoding body: %v", err)
	}

	if body.City != data.City {
		t.Errorf("expected city %v, got %v", data.City, body.City)
	}

	if !body.Temperature.Equal(data.Temperature) {
		t.Errorf("expected temperature %v, got %v", data.Temperature, body.Temperature)
	}
}
