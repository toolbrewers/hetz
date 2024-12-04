package models

import "time"

type Session struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	RoleID    string    `db:"role_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}