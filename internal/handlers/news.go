package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// GetEconomicNews возвращает последние экономические новости.
func GetEconomicNews(w http.ResponseWriter, r *http.Request) {
	news, err := services.GetEconomicNews()
	if err != nil {
		http.Error(w, "Failed to fetch economic news", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
