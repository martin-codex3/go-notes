package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// we will attempt to load the .env file here
type DatabaseConfig struct {
	DatabaseUrl string
	Port string
}


func Load() (*DatabaseConfig, error) {
	var err error = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load the envronment variable")
	}

	var config *DatabaseConfig = &DatabaseConfig{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}

	return config, nil

}