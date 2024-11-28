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
	events, err := services.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get events",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"events":  events,
	})
}

func NewEvent(c *gin.Context) {
	eventTime, _ := time.Parse("2006-01-02 15:04", c.PostForm("event_time"))
	userID, _ := strconv.Atoi(c.PostForm("user_id"))

	event := models.NewEvent(
		c.PostForm("title"),
		c.PostForm("description"),
		c.PostForm("location"),
		eventTime,
		userID,
	)

	err := services.InsertEvent(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to insert events",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
		"event":   event,
	})
}
