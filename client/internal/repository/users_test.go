package repository_test

import (
	"database/sql"
	"testing"

	"hetz-client/internal/models"
	"hetz-client/internal/repository"

	"github.com/doug-martin/goqu/v9"
)

func TestUnit_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    *models.CreateUser
		wantErr error
		cleanup func(db *sql.DB)
	}{
		{
			name: "happy path",
			user: &models.CreateUser{
				Username:     "johndoe",
				Email:        "johndoe12@example.com",
				Password:     "password123",
				HetznerToken: "abc123",
			},
			wantErr: nil,
			cleanup: func(db *sql.DB) {
				query, _, _ := goqu.Delete("users").
					Where(goqu.Ex{"username": "johndoe"}).ToSQL()

				_, err := db.Exec(query)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := repository.New("../../db/app.db")
			defer tc.cleanup(repo.DB)

			_, err := repo.CreateUser(tc.user)
			if err != tc.wantErr {
				t.Errorf("expected: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
