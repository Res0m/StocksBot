package services

import (
	"log"

	"StocksBot/main/internal/external"
)

// SendNotification отправляет уведомление пользователю.
func SendNotification(userID string, text string) error {
	// Логируем попытку отправки уведомления
	log.Printf("Attempting to send notification to user %s: %s", userID, text)

	// Вызываем внешний сервис для отправки уведомления
	err := external.SendNotification(userID, text)
	if err != nil {
		log.Printf("Failed to send notification to user %s: %v", userID, err)
		return err
	}

	// Логируем успешную отправку
	log.Printf("Notification successfully sent to user %s", userID)
	return nil
}
