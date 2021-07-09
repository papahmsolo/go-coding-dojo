package sqlmock

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// TODO:
// implement toDBUser() + DTO* - prosharit'sa
// COALESCE - to read to whom it may concern
// args for prepare()
// common tests for success and error cases
// add error cases (no rows, db error)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	cases := []struct {
		name         string
		expectedUser User
		prepare      func(User)
	}{
		{
			name: "success",
			expectedUser: User{
				ID:             42,
				Email:          "test@gmail.com",
				Password:       "123456",
				ActivationDate: time.Time{},
			},
			prepare: func(u User) {
				mock.ExpectQuery("SELECT (.*) FROM users WHERE id = (.*)").
					WithArgs(u.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "activation_date"}).
						AddRow(u.ID, u.Email, u.Password, u.ActivationDate))
			},
		},
		{
			name: "NULL date",
			expectedUser: User{
				ID:             42,
				Email:          "test@gmail.com",
				Password:       "123456",
				ActivationDate: time.Time{},
			},
			prepare: func(u User) {
				mock.ExpectQuery("SELECT (.*) FROM users WHERE id = (.*)").
					WithArgs(u.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "activation_date"}).
						AddRow(u.ID, u.Email, u.Password, nil))
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepare(tc.expectedUser)

			gotUser, err := GetUser(db, tc.expectedUser.ID)
			fmt.Printf("%+v", gotUser)
			if err != nil {
				t.Errorf("an error '%s' was not expected when getting user by ID", err)
			}

			if tc.expectedUser != gotUser {
				t.Errorf("expected user to be %+v, but got %+v", tc.expectedUser, gotUser)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
