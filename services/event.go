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
		return fmt.Errorf("Error while executing insert statement: %v", err)
	}

	event.ID = id

	return nil
}

func GetEventById(eventId int64) (event *models.Event, err error) {
	query := `SELECT * FROM events WHERE id=?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error while preparing get statement: %v", err)
	}

	defer statement.Close()

	row := statement.QueryRow(eventId)
	event = &models.Event{}
	err = row.Scan(
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

	return event, nil
}

func DeleteEventById(eventId int64) error {
	_, err := GetEventById(eventId)
	if err != nil {
		return err
	}

	query := `DELETE FROM events WHERE id=?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error while preparing delete statement: %v", err)
	}

	defer statement.Close()

	_, err = statement.Exec(eventId)
	if err != nil {
		return fmt.Errorf("Error while executing delete statement: %v", err)
	}

	return nil
}

func UpdateEventById(event *models.Event) error {
	query := `UPDATE events
			SET title=?, description=?, location=?, date_time=?, user_id=?
			WHERE id=?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("Error while preparing update statement: %v", err)
	}

	defer statement.Close()

	_, err = statement.Exec(
		event.Title,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		event.ID,
	)
	if err != nil {
		return fmt.Errorf("Error while executing update statement: %v", err)
	}

	return nil
}
