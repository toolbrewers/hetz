package repository_test

import (
	"hetz/app/models"
	"hetz/app/repository"
	"testing"
)

func TestCreateUser(t *testing.T) {

	user := &models.CreateUser{
		Username:     "johndoe",
		Email:        "johndoe12@example.com",
		Password:     "password123",
		HetznerToken: "abc123",
	}

	repo := repository.New(&repository.Config{DBPath: "../../db/app.db"})
	_, err := repo.CreateUser(user)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a user", err)
	}
}
