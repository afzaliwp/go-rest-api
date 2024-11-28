package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(server *gin.Engine) {
	HandleIndexRoutes(server)
	HandleEventRoutes(server)
}
