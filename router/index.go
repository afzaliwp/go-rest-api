package router

import (
	"github.com/afzaliwp/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleIndexRoutes(server *gin.Engine) {
	server.GET("/", controllers.Index)
}
