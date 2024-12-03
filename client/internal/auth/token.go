package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(randomBytes)
}
