package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// GetCryptoPrices возвращает цены на криптовалюты.
func GetCryptoPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := services.GetCryptoPrices()
	if err != nil {
		http.Error(w, "Failed to fetch crypto prices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}
