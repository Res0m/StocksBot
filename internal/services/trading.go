package services

import "StocksBot/main/internal/external"

// PlaceOrder размещает ордер на бирже.
func PlaceOrder(symbol string, quantity float64, side string) (string, error) {
	return external.PlaceOrder(symbol, quantity, side)
}
