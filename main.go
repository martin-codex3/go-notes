package main

import (
	"context"
	"go-notes/internals/config"
	"go-notes/internals/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// we will load the envronments here 
	// we have to get the databse connection and load it here 
	var err error
	env, err := config.Load()

	if err != nil {
		log.Printf("Failed to load the .env file, %v", err)
	}

	var pool *pgxpool.Pool
	
	ctx := context.Background()


	pool, err = database.DatabaseConnect(env.DatabaseUrl)

	if err != nil {
		log.Printf("Failed to create the database pool, %v", err)
	}

	// we will retry to connect here 
	defer pool.Close()

	// we will attempt to ping the database connection here 
	err = pool.Ping(ctx)

	if err != nil {
		log.Printf("Failed to connect to the database, %v", err)
	}

	log.Printf("Connected successfully")

	// we will create the router here 
	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)

	// sample endpoints 
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Testing the database connection",
			"status": "success",
		})
	})

	// we will then run the server 
	router.Run()
}