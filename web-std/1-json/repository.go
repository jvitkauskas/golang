package main

import (
	"sync"

	"github.com/shopspring/decimal"
)

type Repository struct {
	mu   sync.RWMutex
	data map[string]decimal.Decimal
}

func NewRepository() *Repository {
	return &Repository{
		data: map[string]decimal.Decimal{
			"Vilnius":  decimal.NewFromFloat(8.1),
			"Kaunas":   decimal.NewFromFloat(11.5),
			"Klaipėda": decimal.NewFromFloat(12.75),
		},
	}
}

func (r *Repository) CityExists(city string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.data[city]
	return ok
}

func (r *Repository) GetTemperatureForCity(city string) (decimal.Decimal, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	temperature, ok := r.data[city]
	return temperature, ok
}

func (r *Repository) SetTemperatureForCity(city string, temperature decimal.Decimal) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[city] = temperature
}
