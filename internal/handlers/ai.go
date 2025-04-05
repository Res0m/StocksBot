package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// GetInvestmentAdvice возвращает инвестиционные советы от ИИ.
func GetInvestmentAdvice(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	advice, err := services.GetInvestmentAdvice(req.Prompt)
	if err != nil {
		http.Error(w, "Failed to get advice from AI", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"advice": advice})
}
