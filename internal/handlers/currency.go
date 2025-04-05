package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// GetCurrencyRates возвращает текущие курсы валют.
func GetCurrencyRates(w http.ResponseWriter, r *http.Request) {
	rates, err := services.GetCurrencyRates()
	if err != nil {
		http.Error(w, "Failed to fetch currency rates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}
