package external

import (
	"bytes"
	"fmt"
	"net/http"
)

// MakeRequest выполняет HTTP-запрос и возвращает результат.
func MakeRequest(method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result []byte
	_, err = resp.Body.Read(result)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return result, nil
}
