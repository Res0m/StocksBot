package models

// CryptoPrice представляет модель цены криптовалюты.
type CryptoPrice struct {
	Symbol string  `json:"symbol"` // Символ криптовалюты (например, "BTC", "ETH")
	Price  float64 `json:"price"`  // Цена криптовалюты
}
