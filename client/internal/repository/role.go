package repository

import (
	"hetz-client/internal/models"

	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) CreateRole(role *models.Role) (string, error) {
	query, _, err := goqu.Insert("roles").Rows(
		goqu.Record{
			"name": role.Name,
		},
	).Returning("name").ToSQL()

	if err != nil {
		return "", err
	}

	var name string
	if err := r.DB.QueryRow(query).Scan(&name); err != nil {
		return "", err
	}

	return name, nil
}

func (r *Repository) GetRole(name string) (*models.Role, error) {
	query, _, err := goqu.Select("name").From("roles").Where(
		goqu.C("name").Eq(name),
	).ToSQL()

	if err != nil {
		return nil, err
	}
	var role models.Role
	if err := r.DB.QueryRow(query).Scan(&role.Name); err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *Repository) DeleteRole(name string) error {
	query, _, err := goqu.Delete("roles").Where(
		goqu.C("name").Eq(name),
	).ToSQL()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query)
	return err
}
