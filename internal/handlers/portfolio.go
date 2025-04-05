package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// CreatePortfolio создает новый портфель пользователя.
func CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string   `json:"user_id"`
		Assets []string `json:"assets"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := services.CreatePortfolio(req.UserID, req.Assets)
	if err != nil {
		http.Error(w, "Failed to create portfolio", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Portfolio created successfully"))
}

// GetPortfolio возвращает портфель пользователя.
func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	portfolio, err := services.GetPortfolio(userID)
	if err != nil {
		http.Error(w, "Failed to fetch portfolio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(portfolio)
}
