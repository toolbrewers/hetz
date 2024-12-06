package auth_test

import (
	"testing"
	"time"

	"hetz-client/internal/auth"
)

func TestGenerateToken(t *testing.T) {
	token, err := auth.GenerateToken(time.Hour * 24)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}
	tokenLength := len(token)
	if tokenLength != 64 {
		t.Errorf("Expected token length to be 64, got %d", tokenLength)
	}
}
