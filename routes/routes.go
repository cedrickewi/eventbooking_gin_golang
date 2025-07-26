package routes

import (
	"github.com/cedrickewi/gin_testapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//auth routes
	server.POST("/signup", signup)
	server.POST("/login", login)

	//open routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsByID)

	// authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
