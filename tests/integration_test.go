package tests

import (
	"StocksBot/main/internal/handlers"
	"StocksBot/main/internal/storage"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestIntegrationPortfolio(t *testing.T) {

	db, err := storage.InitDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	
	portfolioStorage := storage.NewPortfolioStorage(db)
	err = portfolioStorage.CreatePortfolio("user123", []string{"BTC", "ETH"})
	if err != nil {
		t.Fatalf("Failed to create portfolio: %v", err)
	}

	req := httptest.NewRequest("GET", "/portfolio?user_id=user123", nil)
	w := httptest.NewRecorder()

	handlers.GetPortfolio(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
