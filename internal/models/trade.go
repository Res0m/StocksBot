package models

// Trade представляет модель торговой операции.
type Trade struct {
	ID       int     `json:"id"`       // Уникальный ID операции
	UserID   string  `json:"user_id"`  // ID пользователя
	Symbol   string  `json:"symbol"`   // Символ актива (например, "AAPL", "BTC")
	Quantity float64 `json:"quantity"` // Количество актива
	Side     string  `json:"side"`     // Направление операции ("buy" или "sell")
}
