package logic

import "github.com/google/uuid"

func GenerateUUID(input interface{}, params ...interface{}) (interface{}, error) {
	return uuid.New().String(), nil
}
