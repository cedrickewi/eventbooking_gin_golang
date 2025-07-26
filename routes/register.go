package routes

import (
	"net/http"
	"strconv"

	"github.com/cedrickewi/gin_testapi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userID")
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	if userId == event.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not register for your own event"})
		return
	}

	if err := event.Register(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userID")
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	if userId == event.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not delete your own event"})
		return
	}

	if err := event.CancelRegistration(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could cancel user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "cancelled successfully"})

}
