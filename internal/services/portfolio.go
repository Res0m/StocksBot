package services

import (
	"StocksBot/main/internal/storage"
	"database/sql"
)

var db *sql.DB // Глобальная переменная для хранения подключения к базе данных

// InitDatabase инициализирует базу данных при старте приложения.
func InitDatabase(dataSourceName string) error {
	var err error
	db, err = storage.InitDB(dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

// CreatePortfolio создает новый портфель пользователя.
func CreatePortfolio(userID string, assets []string) error {
	portfolioStorage := storage.NewPortfolioStorage(db)
	return portfolioStorage.CreatePortfolio(userID, assets)
}

// GetPortfolio получает портфель пользователя.
func GetPortfolio(userID string) (map[string]interface{}, error) {
	portfolioStorage := storage.NewPortfolioStorage(db)
	return portfolioStorage.GetPortfolio(userID)
}
