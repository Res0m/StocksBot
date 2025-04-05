package services

import "StocksBot/main/internal/external"

// GetInvestmentAdvice получает инвестиционные советы от ИИ.
func GetInvestmentAdvice(prompt string) (string, error) {
	return external.GetInvestmentAdvice(prompt)
}
