package routes

import (
	"events.com/m/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticatedAPIs := server.Group("/")
	authenticatedAPIs.Use(middlewares.Authenticate)
	authenticatedAPIs.POST("/events", createEvent)
	authenticatedAPIs.PUT("/events/:id", updateEvent)
	authenticatedAPIs.DELETE("events/:id", deleteEvent)
	authenticatedAPIs.POST("events/:id/register", registerEvent)
	authenticatedAPIs.DELETE("events/:id/register", cancelRegisteration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
