package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {
	connection_string := "user=postgres password=password dbname=reactor_game sslmode=disable"
	db, err := sqlx.Connect("postgres", connection_string)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}
