package models

import (
	"math/rand"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `binding:"required" ,json:"title"`
	Description string    `json:"description"`
	Location    string    `binding:"required" ,json:"location"`
	DateTime    time.Time `binding:"required" ,json:"datetime"`
	UserId      int       `binding:"required" ,json:"userId"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewEvent(title, description string, location string, dateTime time.Time, userId int) Event {
	id := rand.Intn(9000) + 1000
	return Event{
		ID:          id,
		Title:       title,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
		UserId:      userId,
		CreatedAt:   time.Now(),
	}
}

func (e *Event) Save() (success bool, err error) {
	//It should save the event in the database comes from services package

	return true, nil
}
