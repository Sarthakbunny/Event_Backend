package routes

import (
	"events.com/m/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
