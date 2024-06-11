package main

import (
	"log"
	"os"

	"github.com/dexter0323/go-learn/api/db"
	"github.com/dexter0323/go-learn/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
