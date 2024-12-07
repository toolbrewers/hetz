package token_test

import (
	"testing"
	"time"

	token "hetz-client/internal/auth/token"
)

func TestGenerateTokenWithDuration(t *testing.T) {
	token, err := token.GenerateTokenWithDuration(time.Hour * 24)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}
	tokenLength := len(token)
	if tokenLength != 64 {
		t.Errorf("Expected token length to be 64, got %d", tokenLength)
	}
}

func TestGenerateTokenExpiresAt(t *testing.T) {
	expiryTime := time.Now().Add(time.Hour * 24)
	token, err := token.GenerateTokenExpiresAt(expiryTime)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}
	tokenLength := len(token)
	if tokenLength != 64 {
		t.Errorf("Expected token length to be 64, got %d", tokenLength)
	}
}
