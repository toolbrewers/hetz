package repository

import "hetz/app/models"

func (r *Repository) CreateUser(user *models.CreateUser) (uint64, error) {
	query := `INSERT INTO users (username, email, password, hetzner_token)
	VALUES ($1, $2, $3, $4) RETURNING id;`

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
