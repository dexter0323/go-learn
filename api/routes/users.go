package routes

import (
	"fmt"
	"net/http"

	"github.com/dexter0323/go-learn/api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse json request body."})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}
