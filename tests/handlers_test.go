package tests

import (
	"StocksBot/main/internal/handlers"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestGetCurrencyRatesHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/currency-rates", nil)
	w := httptest.NewRecorder()

	handlers.GetCurrencyRates(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
