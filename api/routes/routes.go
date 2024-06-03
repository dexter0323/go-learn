package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

	authRoutes := server.Group("/")
	authRoutes.Use(Authenticate)
	authRoutes.POST("/events", CreateEvent)
	authRoutes.PUT("/events/:id", UpdateEvent)
	authRoutes.DELETE("/events/:id", DeleteEvent)
}
