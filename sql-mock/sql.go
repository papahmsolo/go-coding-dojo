package sqlmock

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func GetUser(db *sql.DB, id int) (User, error) {
	const query = "SELECT id, email, password FROM users WHERE id = ?"
	var u User
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return User{}, fmt.Errorf("cannot get user from DB %w", err)
	}
	return u, nil
}
