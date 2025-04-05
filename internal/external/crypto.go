package external

import (
    "encoding/json"
    "net/http"
)

// FetchCryptoPrices получает цены на криптовалюты.
func FetchCryptoPrices() (map[string]float64, error) {
    resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result map[string]map[string]float64
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    prices := make(map[string]float64)
    for coin, data := range result {
        prices[coin] = data["usd"]
    }

    return prices, nil
}