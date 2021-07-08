package sqlmock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectedUser := User{ID: 42, Email: "test@gmail.com", Password: "123456"}

	mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = (.*)").
		WithArgs(42).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(42, "test@gmail.com", "123456"))

	gotUser, err := GetUser(db, 42)
	if err != nil {
		t.Errorf("an error '%s' was not expected when getting user by ID", err)
	}

	if expectedUser != gotUser {
		t.Errorf("expected user to be %+v, but got %+v", expectedUser, gotUser)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
