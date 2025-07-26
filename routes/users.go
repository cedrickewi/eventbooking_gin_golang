package routes

import (
	"fmt"
	"net/http"

	"github.com/cedrickewi/gin_testapi/models"
	"github.com/cedrickewi/gin_testapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register users. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"erroe": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
