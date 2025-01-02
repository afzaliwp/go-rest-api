package router

import (
	"github.com/afzaliwp/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleAuthRoutes(server *gin.Engine) {
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
}
