package utils

import (
	"github.com/google/uuid"
)

// 生成一个UUID
func GetUUID() uuid.UUID {
	return uuid.New()
}
