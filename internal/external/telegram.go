package external

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendNotification отправляет уведомление через Telegram API.
func SendNotification(userID string, text string) error {
	// Создаем экземпляр Telegram бота
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		return fmt.Errorf("failed to create Telegram bot: %v", err)
	}

	// Получаем chatID пользователя по его ID
	chatID, err := getUserChatID(userID)
	if err != nil {
		return fmt.Errorf("failed to get chatID for user %s: %v", userID, err)
	}

	// Создаем сообщение
	msg := tgbotapi.NewMessage(chatID, text)
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message to user %s: %v", userID, err)
	}

	return nil
}

// getUserChatID получает chatID пользователя по его ID.
func getUserChatID(userID string) (int64, error) {
	// Здесь можно реализовать логику получения chatID из базы данных или кэша
	// Пример:
	switch userID {
	case "user1":
		return 123456789, nil
	case "user2":
		return 987654321, nil
	default:
		return 0, fmt.Errorf("user %s not found", userID)
	}
}
