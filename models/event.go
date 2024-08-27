package models

import (
	"time"

	"events.com/m/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	insertQuery := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	updateQuery := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(updateQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete(id int64) error {
	deleteQuery := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(deleteQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	selectAllEventsQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(selectAllEventsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	searchQuery := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(searchQuery, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
