package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

type Config struct {
	DBPath string
}

var DB *Repository

// aka. default config
func NewConfig() *Config {
	return &Config{
		DBPath: "db/app.db",
	}
}

func New(config *Config) *Repository {
	db, err := sql.Open("sqlite3", config.DBPath)
	if err != nil {
		panic(err)
	}

	return &Repository{DB: db}
}

func (r *Repository) Up() error {
	return r.DB.Ping()
}

func (r *Repository) Close() error {
	return r.DB.Close()
}
