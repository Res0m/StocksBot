package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetInvestmentAdvice(prompt string) (string, error) {
	// Формируем тело запроса
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create request body: %v", err)
	}

	log.Printf("Request to OpenAI: %+v", string(requestBody))

	// Отправляем запрос к OpenAI API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to OpenAI: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode OpenAI response: %v", err)
	}

	log.Printf("Response from OpenAI: %+v", result)

	// Проверяем статус ответа
	if resp.StatusCode != 200 {
		errorMessage, _ := result["error"].(map[string]interface{})
		return "", fmt.Errorf("OpenAI API error: %v", errorMessage["message"])
	}

	// Извлекаем текст ответа
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices in OpenAI response")
	}

	message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid message format in OpenAI response")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("invalid content format in OpenAI response")
	}

	return content, nil
}
