package storage

import (
    "database/sql"
)

// UserStorage предоставляет методы для работы с пользователями.
type UserStorage struct {
    DB *sql.DB
}

// NewUserStorage создает новый экземпляр UserStorage.
func NewUserStorage(db *sql.DB) *UserStorage {
    return &UserStorage{DB: db}
}

// CreateUser создает нового пользователя.
func (s *UserStorage) CreateUser(username string, email string) error {
    _, err := s.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
    return err
}

// GetUser получает пользователя по ID.
func (s *UserStorage) GetUser(userID string) (map[string]interface{}, error) {
    row := s.DB.QueryRow("SELECT username, email FROM users WHERE id = ?", userID)

    var username, email string
    err := row.Scan(&username, &email)
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "id":       userID,
        "username": username,
        "email":    email,
    }, nil
}