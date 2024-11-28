package services

import (
	"fmt"
	"github.com/afzaliwp/go-rest-api/db"
	"github.com/afzaliwp/go-rest-api/models"
)

func GetEvents() ([]models.Event, error) {
	query, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var events []models.Event

	for query.Next() {
		var event models.Event
		err = query.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
			&event.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func InsertEvent(event *models.Event) (err error) {
	query := `INSERT INTO events(title, description, location, date_time, user_id, created_at)
				VALUES(?, ?, ?, ?, ?, ?)
				`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error while preparing statement: %v", err)
	}

	defer statement.Close()

	result, err := statement.Exec(
		event.Title,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		event.CreatedAt,
	)

	id, _ := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("Error while executing statement: %v", err)
	}

	event.ID = id

	return nil
}
