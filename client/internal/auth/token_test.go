package auth_test

import (
	"testing"
	"time"

	"hetz-client/internal/auth"
)

func TestGenerateTokenWithDuration(t *testing.T) {
	token, err := auth.GenerateTokenWithDuration(time.Hour * 24)
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
	token, err := auth.GenerateTokenExpiresAt(expiryTime)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}
	tokenLength := len(token)
	if tokenLength != 64 {
		t.Errorf("Expected token length to be 64, got %d", tokenLength)
	}
}
