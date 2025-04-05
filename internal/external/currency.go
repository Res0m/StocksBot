package external

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func FetchCurrencyRates() (map[string]float64, error) {
	// URL для запроса курсов валют
	resp, err := http.Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currency rates: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode currency rates: %v", err)
	}

	log.Printf("Currency API response: %+v", result)

	// Извлекаем курсы валют
	rates, ok := result["rates"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid rates format in currency API response")
	}

	// Преобразуем курсы в map[string]float64
	ratesMap := make(map[string]float64)
	for currency, rate := range rates {
		ratesMap[currency] = rate.(float64)
	}

	return ratesMap, nil
}
func FetchWeeklyCurrencyHistory(baseCurrency string) (map[string]float64, error) {
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ALPHA_VANTAGE_API_KEY is not set")
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%s&to_symbol=USD&apikey=%s", baseCurrency, apiKey)
	log.Printf("Request to Alpha Vantage: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly currency history: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Alpha Vantage API returned status code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	timeSeries, ok := result["Time Series FX (Daily)"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid time series format in response")
	}

	rates := make(map[string]float64)
	count := 0
	for date, values := range timeSeries {
		if count >= 7 {
			break
		}

		valueMap, ok := values.(map[string]interface{})
		if !ok {
			continue
		}

		closeRate, _ := valueMap["4. close"].(string)
		rate, err := strconv.ParseFloat(closeRate, 64)
		if err != nil {
			continue
		}

		rates[date] = rate
		count++
	}

	return rates, nil
}
func FetchCurrencyHistoryRange(baseCurrency, startDate, endDate string) (map[string]float64, error) {
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ALPHA_VANTAGE_API_KEY is not set")
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=FX_DAILY&from_symbol=%s&to_symbol=USD&apikey=%s", baseCurrency, apiKey)
	log.Printf("Request to Alpha Vantage: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly currency history: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Alpha Vantage API returned status code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	timeSeries, ok := result["Time Series FX (Daily)"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid time series format in response")
	}

	rates := make(map[string]float64)
	for date, values := range timeSeries {
		if date < startDate || date > endDate {
			continue
		}

		valueMap, ok := values.(map[string]interface{})
		if !ok {
			continue
		}

		closeRate, _ := valueMap["4. close"].(string)
		rate, err := strconv.ParseFloat(closeRate, 64)
		if err != nil {
			continue
		}

		rates[date] = rate
	}

	return rates, nil
}


func FetchSupportedCurrencies() ([]string, error) {
    apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("EXCHANGE_RATE_API_KEY is not set")
    }

    url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/USD", apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch supported currencies: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    rates, ok := result["conversion_rates"].(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid conversion_rates format in response")
    }

    currencies := make([]string, 0, len(rates))
    for currency := range rates {
        currencies = append(currencies, currency)
    }

    return currencies, nil
}