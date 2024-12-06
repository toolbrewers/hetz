package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// Takes in a duration and returns a token with the duration encoded in the token
func GenerateTokenWithDuration(tokenDuration time.Duration) (string, error) {
	expiresAt := time.Now().Add(tokenDuration)
	formattedDuration := expiresAt.Format("20060102150405")
	randomBytesLength := 32 - (len(formattedDuration) / 2)

	randomBytes := make([]byte, randomBytesLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := formattedDuration + hex.EncodeToString(randomBytes)
	return token, nil
}

// Takes in expiry time and returns a token with the expiry time encoded in the token
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
