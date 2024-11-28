package router

import (
	"github.com/afzaliwp/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.POST("/event", controllers.NewEvent)
}
