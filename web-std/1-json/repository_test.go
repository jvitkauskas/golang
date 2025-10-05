package main

import (
	"sync"
	"testing"

	"github.com/shopspring/decimal"
)

// newTestRepository creates a new repository with some test data.
func newTestRepository() *Repository {
	return &Repository{
		data: map[string]decimal.Decimal{
			"Vilnius": decimal.NewFromFloat(8.1),
		},
	}
}

func TestCityExists(t *testing.T) {
	repo := newTestRepository()

	t.Run("existing city", func(t *testing.T) {
		if !repo.CityExists("Vilnius") {
			t.Error("expected Vilnius to exist, but it doesn't")
		}
	})

	t.Run("non-existing city", func(t *testing.T) {
		if repo.CityExists("London") {
			t.Error("expected London to not exist, but it does")
		}
	})
}

func TestGetTemperatureForCity(t *testing.T) {
	repo := newTestRepository()

	t.Run("existing city", func(t *testing.T) {
		temp, ok := repo.GetTemperatureForCity("Vilnius")
		if !ok {
			t.Error("expected to get temperature for Vilnius, but didn't")
		}

		expected := decimal.NewFromFloat(8.1)
		if !temp.Equal(expected) {
			t.Errorf("expected temperature %v, got %v", expected, temp)
		}
	})

	t.Run("non-existing city", func(t *testing.T) {
		_, ok := repo.GetTemperatureForCity("London")
		if ok {
			t.Error("expected to not get temperature for London, but did")
		}
	})
}

func TestSetTemperatureForCity(t *testing.T) {
	repo := newTestRepository()

	// Set the temperature for a new city
	city := "London"
	temp := decimal.NewFromFloat(15.5)
	repo.SetTemperatureForCity(city, temp)

	// Check if the city exists
	if !repo.CityExists(city) {
		t.Errorf("expected %v to exist, but it doesn't", city)
	}

	// Check if the temperature is correct
	newTemp, ok := repo.GetTemperatureForCity(city)
	if !ok {
		t.Errorf("expected to get temperature for %v, but didn't", city)
	}

	if !newTemp.Equal(temp) {
		t.Errorf("expected temperature %v, got %v", temp, newTemp)
	}
}

func TestRepositoryConcurrentAccess(t *testing.T) {
	repo := newTestRepository()

	// This test will run a bunch of goroutines to read and write to the repository concurrently.
	// If the repository is not thread-safe, this test will likely fail with a race condition.

	var wg sync.WaitGroup

	// Number of goroutines to spawn
	numGoroutines := 100

	// Number of iterations per goroutine
	numIterations := 100

	// Spawn a bunch of goroutines to read and write
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < numIterations; j++ {
				// Read
				repo.GetTemperatureForCity("Vilnius")

				// Write
				repo.SetTemperatureForCity("London", decimal.NewFromFloat(15.5))
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
