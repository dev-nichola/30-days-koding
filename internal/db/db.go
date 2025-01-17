package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewDB() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load Env")
	}

	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal("Connection error", err)
	}

	if err := db.Ping(); err != nil {
		panic("Failed to connect db")
	}

	return db, nil
}
