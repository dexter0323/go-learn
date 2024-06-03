package main

import (
	"net/http"
	"strconv"

	"github.com/dexter0323/go-learn/api/db"
	"github.com/dexter0323/go-learn/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()

	server.GET("/events", func(ctx *gin.Context) {
		events, err := models.GetEvents()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
			return
		}
		ctx.JSON(http.StatusOK, events)
	})

	server.GET("/events/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
			return
		}

		event, err := models.GetEvent(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
			return
		}
		ctx.JSON(http.StatusOK, event)
	})

	server.POST("/events", func(ctx *gin.Context) {
		var event models.Event
		err := ctx.ShouldBindJSON(&event)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse json request body."})
			return
		}

		event.ID = 1
		event.UserId = 1

		err = event.Save()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"event": event})
	})

	server.Run(":8080")
}
