package services

import "StocksBot/main/internal/external"

// GetCryptoPrices получает цены на криптовалюты.
func GetCryptoPrices() (map[string]float64, error) {
	return external.FetchCryptoPrices()
}
