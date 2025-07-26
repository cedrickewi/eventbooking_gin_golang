package models

import (
	"time"

	"github.com/cedrickewi/gin_testapi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (event *Event) Save() error {
	//Later: add it to database
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		 VALUES(?, ?, ?, ?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	var events = []Event{}

	query := `
		SELECT id , name, description, location, dateTime, user_id FROM events 
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event

	query := `
		SELECT * FROM events WHERE id=?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	if err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}

func Update(event Event) error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id int64) error {
	query := `
		DELETE FROM events WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
