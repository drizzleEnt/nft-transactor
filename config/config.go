package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// PostgreSQL
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int

	HTTP_HOST string
	HTTP_PORT string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed load env file: %s", err.Error())
	}

	DB_USER = os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Fatalf("DB_USER not set, but required")
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Fatalf("DB_PASSWORD not set, but required")
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Fatalf("DB_NAME not set, but required")
	}

	DB_HOST = os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Fatalf("DB_HOST not set, but required")
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to a number: %v", err)
	}
	DB_PORT = dbPort

	HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		log.Fatalf("HTTP_HOST not set, but required")
	}

	HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		log.Fatalf("HTTP_PORT not set, but required")
	}
}
