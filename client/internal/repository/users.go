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

func (r *Repository) GetUserByEmail(email string) (*models.GetUser, error) {
	query, _, err := goqu.Select("id", "email", "username").From("users").Where(goqu.C("email").Eq(email)).ToSQL()
	if err != nil {
		return nil, err
	}

	var user models.GetUser
	if err := r.DB.QueryRow(query).Scan(&user.ID, &user.Email, &user.Username); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserPassword(email string) (*models.GetUserPassword, error) {
	query, _, err := goqu.Select("password").From("users").Where(goqu.C("email").Eq(email)).ToSQL()
	if err != nil {
		return nil, err
	}

	var user models.GetUserPassword
	if err := r.DB.QueryRow(query).Scan(&user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUserByID(id uint64) error {
	query, _, err := goqu.Delete("users").Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(query)
	return err
}
