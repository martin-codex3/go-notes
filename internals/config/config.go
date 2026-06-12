package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// we first lead the envroment variables here
type DatabaseEnvronment struct {
	DatabaseUrl string
	Port string
}

// the function to load the envronment variables here
func Load() (*DatabaseEnvronment, error) {
	
	// we will load the envronment variables here 
	var err error = godotenv.Load()

	if err != nil {
		log.Printf("Failed to load the envronment variables")
	}

	// we will have to pass the variables to the config variable
	var config *DatabaseEnvronment = &DatabaseEnvronment{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}

	return config, nil
}