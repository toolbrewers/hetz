package repository_test

import (
	"testing"

	"hetz-client/internal/models"
	"hetz-client/internal/repository"
)

func TestUser(t *testing.T) {
	t.Run("CreateUser", testCreateUser)
	t.Run("GetUserByEmail", testGetUserByEmail)
	t.Run("DeleteUserByID", testDeleteUser)
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

func testGetUserByEmail(t *testing.T) {
	repo := repository.New("../../db/client.db")
	defer repo.Close()

	user, err := repo.GetUserByEmail("johndoe12@example.com")
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
	if user.Email != "johndoe12@example.com" {
		t.Errorf("expected: %v, got: %v", "johndoe12@example.com", user.Email)
	}
}

func testDeleteUser(t *testing.T) {
	repo := repository.New("../../db/client.db")
	defer repo.Close()

	user, err := repo.GetUserByEmail("johndoe12@example.com")
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}

	err = repo.DeleteUserByID(user.ID)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}
}
