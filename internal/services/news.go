package services

import (
	"StocksBot/main/internal/external"
)

func GetEconomicNews() ([]string, error) {
	news, err := external.FetchEconomicNews()
	if err != nil {
		return nil, err
	}

	// Преобразуем массив словарей в массив строк
	var newsTitles []string
	for _, item := range news {
		title := item["title"]
		if title != "" {
			newsTitles = append(newsTitles, title)
		}
	}

	return newsTitles, nil
}
