package models

import "time"

type Session struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	UserAgent string    `db:"user_agent"`
	IPAddress string    `db:"ip_address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Just the bare basics what's needed to validate a session.
type SessionToken struct {
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}
