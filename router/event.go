package router

import (
	"github.com/afzaliwp/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventById)

	server.POST("/event", controllers.NewEvent)
	//server.POST("/event/:id/register", controllers.RegisterUserOnEvent)

	server.PUT("/event/:id", controllers.UpdateEvent)

	server.DELETE("/event/:id", controllers.DeleteEvent)
	//server.DELETE("/event/:id/register", controllers.DeleteUserFromEvent)
}
