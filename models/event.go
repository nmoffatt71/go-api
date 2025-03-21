package models

import (
	"fmt"
	"time"

	"rest-api.com/m/v2/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	// Add to DB
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT *
	FROM events 
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `
	SELECT *
	FROM events 
	WHERE id = ?
	`

	row := db.DB.QueryRow(query, id)
	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &event, nil

}
func (e Event) UpdateEventByID() error {
	// Add to DB
	query := `
	UPDATE events 
	SET 
	name = ?, 
	description = ?,
	location = ?, 
	dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err

}

func (e Event) DeleteEventByID() error {
	// Add to DB
	query := `
	DELETE FROM events 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err

}
