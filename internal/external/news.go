package external

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func FetchEconomicNews() ([]map[string]string, error) {
	// URL для запроса новостей
	apiKey := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?category=business&apiKey=%s", apiKey)

	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode news response: %v", err)
	}

	log.Printf("News API response: %+v", result)

	// Проверяем статус ответа
	status, ok := result["status"].(string)
	if !ok || status != "ok" {
		return nil, fmt.Errorf("invalid news API response: %v", result)
	}

	// Извлекаем статьи
	articles, ok := result["articles"].([]interface{})
	if !ok || len(articles) == 0 {
		return nil, fmt.Errorf("no articles found in news API response")
	}

	// Формируем список новостей
	var news []map[string]string
	for _, article := range articles {
		articleMap, ok := article.(map[string]interface{})
		if !ok {
			continue
		}

		title, _ := articleMap["title"].(string)
		url, _ := articleMap["url"].(string)

		news = append(news, map[string]string{
			"title": title,
			"url":   url,
		})
	}

	return news, nil
}
