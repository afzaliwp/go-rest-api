package models

import (
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `binding:"required" ,json:"title"`
	Description string    `json:"description"`
	Location    string    `binding:"required" ,json:"location"`
	DateTime    time.Time `binding:"required" ,json:"datetime"`
	UserId      int64     `binding:"required" ,json:"userId"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewEvent(title, description string, location string, dateTime time.Time, userId int) Event {
	return Event{
		Title:       title,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
		UserId:      int64(userId),
		CreatedAt:   time.Now(),
	}
}
