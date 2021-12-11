package utility

import (
	"github.com/google/uuid"
)

func ConvertUUID(s string) uuid.UUID {
	res, _ := uuid.Parse(s)
	return res
}
