package models

// Notification представляет модель уведомления.
type Notification struct {
	ID     int    `json:"id"`      // Уникальный ID уведомления
	UserID string `json:"user_id"` // ID пользователя
	Text   string `json:"text"`    // Текст уведомления
}
