package utils

import "github.com/google/uuid"

func GetNextId() string {
	return uuid.New().String()
}
