package models

import (
	"example.com/event-registration-app/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
    INSERT INTO events(name, description, location, dateTime, user_id) 
    VALUES ($1, $2, $3, $4, $5) 
    RETURNING id`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}

	defer stmt.Close()
	// Execute the statement and retrieve the inserted ID
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
