package config

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("dbUrl is required")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}
