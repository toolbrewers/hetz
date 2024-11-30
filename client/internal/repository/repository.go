package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

func New(source string) *Repository {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		panic(fmt.Errorf("failed to open db connection: %w", err))
	}

	return &Repository{DB: db}
}

func (r *Repository) Close() error {
	return r.DB.Close()
}

func (r *Repository) Up() error {
	return r.DB.Ping()
}
