package getid

import "github.com/google/uuid"

// GetUUID returns RFC 4122 version 4 UUID string.
func GetUUID() string {
	return uuid.New().String()
}
