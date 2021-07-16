package sqlmock

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var ErrNotFound = errors.New("no such user")

type User struct {
	ID             int
	Email          string
	Password       string
	ActivationDate time.Time
}

type dbUser struct {
	ID             int
	Email          string
	Password       string
	ActivationDate sql.NullTime
}

func (dbu dbUser) toUser() User {
	return User{
		ID:             dbu.ID,
		Email:          dbu.Email,
		Password:       dbu.Password,
		ActivationDate: dbu.ActivationDate.Time,
	}
}

func GetUser(db *sql.DB, userID int) (User, error) {
	const query = "SELECT id, email, password, activation_date FROM users WHERE id = ?"
	var dbu dbUser
	err := db.QueryRow(query, userID).Scan(
		&dbu.ID,
		&dbu.Email,
		&dbu.Password,
		&dbu.ActivationDate,
	)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("%w: %d", ErrNotFound, userID)
	}
	if err != nil {
		return User{}, fmt.Errorf("cannot get user from DB %w", err)
	}
	return dbu.toUser(), nil
}

func UpdateUser(db *sql.DB, u User) error {
	const query = "UPDATE users SET email=?, password=? WHERE id = ?"
	res, err := db.Exec(query, u.Email, u.Password, u.ID)
	if err != nil {
		return fmt.Errorf("cannot update user in DB: %w", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot get RowsAffected result: %w", err)
	}
	if n == 0 {
		return fmt.Errorf("user was not updated. %w: %d", ErrNotFound, u.ID)
	}
	return nil
}

func InsertUser(db *sql.DB, u User) error {
	const query = "INSERT INTO users (email, password, activation_date) VALUES(?, ?, ?)"
	_, err := db.Exec(query, u.Email, u.Password, u.ActivationDate)
	if err != nil {
		return fmt.Errorf("cannot insert user into DB: %w", err)
	}
	return nil
}
