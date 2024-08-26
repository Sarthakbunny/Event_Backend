package models

import "events.com/m/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	saveQuery := `INSERT INTO users (email, password)
	VALUES (?, ?)`
	stmt, err := db.DB.Prepare(saveQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	_ = user.ID
	return err
}
