package repository_test

import (
	"hetz-client/internal/models"
	"testing"
	"time"
)

var testToken string

func TestSession(t *testing.T) {

	t.Run("CreateSession", testCreateSession)
	t.Run("GetSessionToken", testGetSessionToken)
	t.Run("DeleteSession", testDeleteSession)
}

func testCreateSession(t *testing.T) {
	repo := openRepo()
	defer repo.Close()

	err := repo.CreateSession(&models.Session{
		Token:     "test-token",
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		t.Errorf("Error creating session: %v", err)
	}
}

func testGetSessionToken(t *testing.T) {
	repo := openRepo()
	defer repo.Close()

	session, err := repo.GetSessionToken("test-token")
	if err != nil {
		t.Errorf("Error getting session: %v", err)
	}
	if session == nil {
		t.Errorf("Session not found")
	}
	if session.Token != "test-token" {
		t.Errorf("Session token mismatch: %v", session.Token)
	}
}

func testDeleteSession(t *testing.T) {
	repo := openRepo()
	defer repo.Close()

	err := repo.DeleteSessionByToken("test-token")
	if err != nil {
		t.Errorf("Error deleting session: %v", err)
	}
}
