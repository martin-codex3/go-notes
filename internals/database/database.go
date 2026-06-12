package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// we will attempt to create the connection with the database here
func DatabaseConnect (DatabaseUrl string) (*pgxpool.Pool, error) {
	// for us to create the database connection we need a context 
	var ctx context.Context = context.Background()

	var err error
	var cfg *pgxpool.Config

	// we have to get the 
	cfg, err = pgxpool.ParseConfig(DatabaseUrl)

	if err != nil {
		log.Fatal("Failed to parse the configuration file")
		return nil, err
	}

	// we will create the database pooling here 
	var pool *pgxpool.Pool

	pool, err = pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		log.Printf("Failed to create the database pool")
		return nil, err
	}

	// we will then have to return the pool here 
}