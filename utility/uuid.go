package utility

import (
	"github.com/google/uuid"
)

func NewUUID() uuid.UUID {
	return uuid.New()
}

func ConvertUUID(s string) uuid.UUID {
	res, _ := uuid.Parse(s)
	return res
}

func ConvertString(uuid uuid.UUID) string {
	return uuid.String()
}
