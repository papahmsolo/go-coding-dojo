package sqlmock

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
	if err != nil {
		return User{}, fmt.Errorf("cannot get user from DB %w", err)
	}
	return dbu.toUser(), nil
}
