package models

// User представляет модель пользователя.
type User struct {
	ID       int    `json:"id"`       // Уникальный ID пользователя
	Username string `json:"username"` // Имя пользователя
	Email    string `json:"email"`    // Email пользователя
}
