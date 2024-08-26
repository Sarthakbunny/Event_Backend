package routes

import (
	"net/http"

	"events.com/m/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "inputs are not valid"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "user": user})
}
