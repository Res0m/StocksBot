package handlers

import (
	"StocksBot/main/internal/services"
	"encoding/json"
	"net/http"
)

// SendNotification отправляет уведомление пользователю.
func SendNotification(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := services.SendNotification(req.UserID, req.Text)
	if err != nil {
		http.Error(w, "Failed to send notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification sent successfully"))
}
