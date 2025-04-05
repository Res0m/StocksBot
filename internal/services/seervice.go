package services

import "fmt"

// BaseService предоставляет общие методы для всех сервисов.
type BaseService struct{}

// LogError логирует ошибки.
func (b *BaseService) LogError(err error, msg string) {
    if err != nil {
        fmt.Printf("Error: %s - %v\n", msg, err)
    }
}