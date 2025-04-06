package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {
	// connection_string := "user=postgres password=password dbname=reactor_game sslmode=disable"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	const maxAttempts = 10
	const delay = 2 * time.Second

	var db *sqlx.DB
	var err error

	for i := 1; i <= maxAttempts; i++ {
		db, err = sqlx.Connect("postgres", connectionString)
		if err == nil {
			// Успешное подключение
			return db, nil
		}

		log.Printf("Попытка %d: не удалось подключиться к БД: %v", i, err)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("не удалось подключиться к базе данных после %d попыток: %v", maxAttempts, err)
}
