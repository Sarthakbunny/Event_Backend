package routes

import (
	"net/http"
	"strconv"

	"events.com/m/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse the id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "fetched successfully",
		"event":   event,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) //similar to scanf

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was an error parsing context data"})
		return
	}

	event.UserId = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was an error saving the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
