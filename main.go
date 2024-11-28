package main

import (
	"fmt"
	"github.com/afzaliwp/go-rest-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	router.InitializeRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		_ = fmt.Errorf("error starting server: %s", err)
		return
	}
}
