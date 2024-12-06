package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func GenerateTokenWithDuration(tokenDuration time.Duration) (string, error) {

	formattedDuration := time.Now().Add(tokenDuration).Format("20060102150405")
	randomBytesLength := 32 - (len(formattedDuration) / 2)

	randomBytes := make([]byte, randomBytesLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := formattedDuration + hex.EncodeToString(randomBytes)
	return token, nil
}

func GenerateTokenExpiresAt(tokenExpiresAt time.Time) (string, error) {
	formattedDuration := tokenExpiresAt.Format("20060102150405")
	randomBytesLength := 32 - (len(formattedDuration) / 2)

	randomBytes := make([]byte, randomBytesLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := formattedDuration + hex.EncodeToString(randomBytes)
	return token, nil
}
