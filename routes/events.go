package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cedrickewi/gin_testapi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventsByID(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	event, err := models.GetEventByID(idInt)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	userID := context.GetInt64("userID")

	event.UserID = userID

	if err := event.Save(); err != nil {
		fmt.Println("error::", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error creating event. Try again later!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(idInt)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}

	if userID != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	updatedEvent.ID = idInt
	if err := models.Update(updatedEvent); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(idInt)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}

	if userID != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to delete event"})
		return
	}

	if err := models.Delete(idInt); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "event deleted successfully"})

}
