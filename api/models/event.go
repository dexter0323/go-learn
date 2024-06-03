package models

import (
	"time"

	"github.com/dexter0323/go-learn/api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {

	query, err := db.DB.Prepare(`INSERT INTO events(name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer query.Close()

	result, err := query.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

func GetEvents() ([]Event, error) {
	query, err := db.DB.Query(`SELECT * FROM events`)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	var events []Event

	for query.Next() {
		var event Event
		err := query.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := db.DB.QueryRow(`SELECT * FROM events WHERE id = ?`, id)
	var event Event

	err := query.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	_, err := db.DB.Exec(`
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	_, err := db.DB.Exec(`DELETE FROM events WHERE id = ?`, e.ID)
	if err != nil {
		return err
	}

	return nil
}
