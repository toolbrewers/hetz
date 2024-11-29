package repository

import (
	"hetz-client/internal/models"

	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) CreateUser(user *models.CreateUser) (uint64, error) {
	query, _, _ := goqu.Insert("users").Rows(
		goqu.Record{
			"username":      user.Username,
			"email":         user.Email,
			"password":      user.Password,
			"hetzner_token": user.HetznerToken,
		},
	).Returning("id").ToSQL()

	var id uint64
	if err := r.DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.HetznerToken,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
