package repository

import (
	"hetz-client/internal/models"

	"github.com/doug-martin/goqu/v9"
)

// Make sure that the role exists before creating a session.
func (r *Repository) CreateSession(session *models.Session) error {
	query, _, err := goqu.Insert("sessions").Rows(session).ToSQL()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query)
	return err
}

func (r *Repository) GetSessionToken(token string) (*models.SessionToken, error) {
	// get token and expires_at from session
	query, _, err := goqu.Select("token", "expires_at").From("sessions").Where(goqu.C("token").Eq(token)).ToSQL()

	if err != nil {
		return nil, err
	}

	var session models.SessionToken

	if err := r.DB.QueryRow(query).Scan(&session.Token, &session.ExpiresAt); err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) DeleteSessionByToken(token string) error {
	query, _, err := goqu.Delete("sessions").Where(goqu.C("token").Eq(token)).ToSQL()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query)
	return err
}

func (r *Repository) DeleteAllSessionsByUserID(userID uint64) error {
	query, _, err := goqu.Delete("sessions").Where(goqu.C("user_id").Eq(userID)).ToSQL()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query)
	return err
}
