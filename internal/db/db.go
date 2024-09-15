package db

import (
	"log"

	"github.com/joho/godotenv"
)

func NewDB() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load Env")
	}

	var()

}
