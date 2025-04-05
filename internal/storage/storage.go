package storage

import "fmt"

// StorageError представляет ошибку хранилища.
type StorageError struct {
	Message string
}

func (e *StorageError) Error() string {
	return fmt.Sprintf("Storage error: %s", e.Message)
}

// CheckError проверяет ошибку и возвращает StorageError, если она есть.
func CheckError(err error, message string) error {
	if err != nil {
		return &StorageError{Message: message}
	}
	return nil
}
