package storage

import (
    "database/sql"
)

// NotificationStorage предоставляет методы для работы с уведомлениями.
type NotificationStorage struct {
    DB *sql.DB
}

// NewNotificationStorage создает новый экземпляр NotificationStorage.
func NewNotificationStorage(db *sql.DB) *NotificationStorage {
    return &NotificationStorage{DB: db}
}

// CreateNotification создает новое уведомление.
func (s *NotificationStorage) CreateNotification(userID string, text string) error {
    _, err := s.DB.Exec("INSERT INTO notifications (user_id, text) VALUES (?, ?)", userID, text)
    return err
}

// GetNotifications получает уведомления пользователя.
func (s *NotificationStorage) GetNotifications(userID string) ([]map[string]interface{}, error) {
    rows, err := s.DB.Query("SELECT id, text FROM notifications WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notifications []map[string]interface{}
    for rows.Next() {
        var id int
        var text string
        err := rows.Scan(&id, &text)
        if err != nil {
            return nil, err
        }

        notifications = append(notifications, map[string]interface{}{
            "id":   id,
            "text": text,
        })
    }

    return notifications, nil
}