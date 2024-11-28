package controllers

import (
	"github.com/afzaliwp/go-rest-api/models"
	"github.com/afzaliwp/go-rest-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"events":  services.GetEvents(),
	})
}

func NewEvent(c *gin.Context) {
	eventTime, _ := time.Parse("2006-01-02 15:04", c.PostForm("event_time"))
	userID, _ := strconv.Atoi(c.PostForm("user_id"))

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
		"event": models.NewEvent(
			c.PostForm("title"),
			c.PostForm("description"),
			c.PostForm("location"),
			eventTime,
			userID,
		),
	})
}
