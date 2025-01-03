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

func GetEventById(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be integer",
			"error":   err.Error(),
		})
	}
	id := int64(idInt)

	event, err := services.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event: " + idParam,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"event":   event,
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

func UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an integer",
			"error":   err.Error(),
		})
	}
	id := int64(idInt)

	event, err := services.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get event: " + idParam,
		})
	}

	// Helper function to update fields
	updateField := func(field *string, formValue string) {
		if formValue != "" {
			*field = formValue
		}
	}

	updateField(&event.Title, c.PostForm("title"))
	updateField(&event.Description, c.PostForm("description"))
	updateField(&event.Location, c.PostForm("location"))

	if eventTimeStr := c.PostForm("event_time"); "" != eventTimeStr {
		eventTime, _ := time.Parse("2006-01-02 15:04", eventTimeStr)
		event.DateTime = eventTime
	}

	if userIdStr := c.PostForm("user_id"); "" != userIdStr {
		userID, _ := strconv.Atoi(userIdStr)
		event.UserId = int64(userID)
	}

	err = services.UpdateEventById(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update event: " + idParam,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"event":   event,
	})
}

func DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an integer",
			"error":   err.Error(),
		})
	}

	id := int64(idInt)
	err = services.DeleteEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete event: " + idParam,
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"event_id": id,
	})
}
