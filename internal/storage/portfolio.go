package storage

import (
    "database/sql"
)

// PortfolioStorage предоставляет методы для работы с портфелем.
type PortfolioStorage struct {
    DB *sql.DB
}

// NewPortfolioStorage создает новый экземпляр PortfolioStorage.
func NewPortfolioStorage(db *sql.DB) *PortfolioStorage {
    return &PortfolioStorage{DB: db}
}

// CreatePortfolio создает новый портфель пользователя.
func (s *PortfolioStorage) CreatePortfolio(userID string, assets []string) error {
    _, err := s.DB.Exec("INSERT INTO portfolios (user_id, assets) VALUES (?, ?)", userID, assets)
    return err
}

// GetPortfolio получает портфель пользователя.
func (s *PortfolioStorage) GetPortfolio(userID string) (map[string]interface{}, error) {
    row := s.DB.QueryRow("SELECT assets FROM portfolios WHERE user_id = ?", userID)

    var assets []byte
    err := row.Scan(&assets)
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "user_id": userID,
        "assets":  string(assets),
    }, nil
}