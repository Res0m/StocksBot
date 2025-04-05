package services

import "StocksBot/main/internal/external"

// GetCurrencyRates получает текущие курсы валют.
func GetCurrencyRates() (map[string]float64, error) {
	return external.FetchCurrencyRates()
}

func GetWeeklyCurrencyHistory(baseCurrency string) (map[string]float64, error) {
	return external.FetchWeeklyCurrencyHistory(baseCurrency)
}

func GetCurrencyHistoryRange(baseCurrency, startDate, endDate string) (map[string]float64, error) {
	return external.FetchCurrencyHistoryRange(baseCurrency, startDate, endDate)
}
func GetSupportedCurrencies() ([]string, error) {
    return external.FetchSupportedCurrencies()
}