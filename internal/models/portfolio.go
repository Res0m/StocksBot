package models

// Portfolio представляет модель портфеля пользователя.
type Portfolio struct {
	ID     int      `json:"id"`      // Уникальный ID портфеля
	UserID string   `json:"user_id"` // ID пользователя
	Assets []string `json:"assets"`  // Список активов в портфеле
}
