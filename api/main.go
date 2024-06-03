package main

import (
	"github.com/dexter0323/go-learn/api/db"
	"github.com/dexter0323/go-learn/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
