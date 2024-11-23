package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

func New() (*Repository, error) {
	db, err := sql.Open("sqlite3", "./storage/app.db")
	if err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}

func (r *Repository) Up() error {
	return r.DB.Ping()
}
