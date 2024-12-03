package auth_test

import (
	"testing"

	"hetz-client/internal/auth"
)

func TestGenerateToken(t *testing.T) {
	token := auth.GenerateToken()
	tokenLength := len(token)
	if tokenLength != 64 {
		t.Errorf("Expected token length to be 64, got %d", tokenLength)
	}
}
