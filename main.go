package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	
	var router gin.Engine = *gin.Default()

	// we will use the router here 
	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This is the test for the server",
		})
	})

	router.Run()
}