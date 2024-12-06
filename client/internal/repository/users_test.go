package repository_test

import (
	"testing"

	"hetz-client/internal/models"
	"hetz-client/internal/repository"
)

func TestUser(t *testing.T) {
	t.Run("CreateUser", testCreateUser)
}

func testCreateUser(t *testing.T) {
	repo := repository.New("../../db/client.db")
	defer repo.Close()

	user := &models.CreateUser{
		Username:     "johndoe",
		Email:        "johndoe12@example.com",
		Password:     "password123",
		HetznerToken: "abc123",
	}
	_, err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
}
