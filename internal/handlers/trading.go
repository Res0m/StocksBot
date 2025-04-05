package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// PlaceOrder размещает ордер на бирже.
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Symbol   string  `json:"symbol"`
		Quantity float64 `json:"quantity"`
		Side     string  `json:"side"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	orderID, err := services.PlaceOrder(req.Symbol, req.Quantity, req.Side)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"order_id": orderID})
}
