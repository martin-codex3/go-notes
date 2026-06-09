package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	DatabaseUrl string
	Port string
}

func Load() (*Config, error) {

	var err error = godotenv.Load()

	if err != nil {
		log.Println("Failed to load the envronment variables")
	}

	config := &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}

	return config, nil
}