package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var conn *sql.DB

func Connect() (*sql.DB, error) {
	er := godotenv.Load()
	if er != nil {
		log.Println("Error loading .env file")
	}

	var err error
	connStr := os.Getenv("DATABASE_URL")
	conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	return conn, nil
}

func GetDB() *sql.DB {
	return conn
}
