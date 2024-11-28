package services

import (
	"github.com/afzaliwp/go-rest-api/models"
	"time"
)

func GetEvents() []models.Event {
	eventTime1, _ := time.Parse("2006-01-02 15:04", "2025-05-12 12:30")
	eventTime2, _ := time.Parse("2006-01-02 15:04", "2027-09-22 13:45")
	events := []models.Event{
		models.NewEvent(
			"title here",
			"description here",
			"France, Paris",
			eventTime1,
			12,
		),

		models.NewEvent(
			"title here 2",
			"description here 2",
			"Iran, Tehran",
			eventTime2,
			5,
		),
	}

	return events
}
