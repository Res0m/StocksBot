package external

import (
    "bytes"
    "encoding/json"
    "net/http"
)

// PlaceOrder отправляет ордер на биржу.
func PlaceOrder(symbol string, quantity float64, side string) (string, error) {
    requestBody := map[string]interface{}{
        "symbol":   symbol,
        "quantity": quantity,
        "side":     side,
    }

    jsonBody, _ := json.Marshal(requestBody)
    resp, err := http.Post(
        "https://api.tradingplatform.com/orders",
        "application/json",
        bytes.NewBuffer(jsonBody),
    )
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    return result["order_id"].(string), nil
}