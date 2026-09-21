package util

import "uuid"

// GetUuid 获取UUID
func GetUuid() string {
	return uuid.NewV7().String()
}
