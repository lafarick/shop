package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Connection struct {
	DB *sql.DB
}

func Connect() (*sql.DB, error) {
	connStr := "host=127.0.0.1 port=6111 user=test password=test dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
