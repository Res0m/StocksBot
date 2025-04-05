package models

// CurrencyRate представляет модель курса валюты.
type CurrencyRate struct {
	Currency string  `json:"currency"` // Код валюты (например, "USD", "EUR")
	Rate     float64 `json:"rate"`     // Обменный курс
}
