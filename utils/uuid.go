package utils

import "github.com/satori/go.uuid"

// 生成一个UUID
func GetUUID() uuid.UUID {
	return uuid.NewV4()
}
