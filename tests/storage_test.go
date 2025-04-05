package tests

import (
	"StocksBot/main/internal/storage"
	"testing"
)

func TestCreatePortfolio(t *testing.T) {
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

	// Проверяем, что портфель создан
	row := db.QueryRow("SELECT assets FROM portfolios WHERE user_id = ?", "user123")
	var assets string
	err = row.Scan(&assets)
	if err != nil {
		t.Fatalf("Failed to fetch portfolio: %v", err)
	}

	if assets != "[\"BTC\",\"ETH\"]" {
		t.Errorf("Expected assets to be [\"BTC\",\"ETH\"], got %s", assets)
	}
}
