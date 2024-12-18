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

	CONTRACT_ADDRESS string
	RPC              string
	TOTAL_SUPPLY_ABI string
	MINT_ABI         string
	PRIVATE_KEY      string
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

	CONTRACT_ADDRESS = os.Getenv("CONTRACT_ADDRESS")
	if CONTRACT_ADDRESS == "" {
		log.Fatalf("CONTRACT_ADDRESS not set, but required")
	}

	RPC = os.Getenv("RPC")
	if RPC == "" {
		log.Fatalf("RPC not set, but required")
	}

	TOTAL_SUPPLY_ABI = os.Getenv("TOTAL_SUPPLY_ABI")
	if TOTAL_SUPPLY_ABI == "" {
		log.Fatalf("TOTAL_SUPPLY_ABI not set, but required")
	}

	MINT_ABI = os.Getenv("MINT_ABI")
	if MINT_ABI == "" {
		log.Fatalf("MINT_ABI not set, but required")
	}

	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	if PRIVATE_KEY == "" {
		log.Fatalf("PRIVATE_KEY not set, but required")
	}
}
