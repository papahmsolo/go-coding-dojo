package sqlmock

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// TODO:
// implement toDBUser() + DTO* - prosharit'sa
// COALESCE - to read to whom it may concern
// args for prepare()
// not catching sql syntax error - write concrete queries with regex

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	var testErr = errors.New("test error")

	cases := []struct {
		name          string
		expectedUser  User
		expectedError error
		prepare       func(User)
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
		{
			name:          "user not found",
			expectedUser:  User{},
			expectedError: ErrNotFound,
			prepare: func(u User) {
				mock.ExpectQuery("SELECT (.*) FROM users WHERE id = (.*)").
					WithArgs(u.ID).
					WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name:          "db call error",
			expectedUser:  User{},
			expectedError: testErr,
			prepare: func(u User) {
				mock.ExpectQuery("SELECT (.*) FROM users WHERE id = (.*)").
					WithArgs(u.ID).
					WillReturnError(testErr)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepare(tc.expectedUser)

			gotUser, err := GetUser(db, tc.expectedUser.ID)
			//fmt.Printf("%+v", gotUser)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error to be %v, but got %v", tc.expectedError, err)
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

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	var testErr = errors.New("test error")

	cases := []struct {
		name          string
		args          User
		expectedError error
		prepare       func(User)
	}{
		{
			name: "success",
			prepare: func(u User) {
				mock.ExpectExec("UPDATE users SET (.*) WHERE id = ?").
					WithArgs(u.Email, u.Password, u.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:          "no such user",
			expectedError: ErrNotFound,
			prepare: func(u User) {
				mock.ExpectExec("UPDATE users SET (.*) WHERE id = ?").
					WithArgs(u.Email, u.Password, u.ID).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name:          "db call error",
			expectedError: testErr,
			prepare: func(u User) {
				mock.ExpectExec("UPDATE users SET (.*) WHERE id = ?").
					WithArgs(u.Email, u.Password, u.ID).
					WillReturnError(testErr)
			},
		},
		{
			name:          "RowsAffected() error",
			expectedError: testErr,
			prepare: func(u User) {
				mock.ExpectExec("UPDATE users SET (.*) WHERE id = ?").
					WithArgs(u.Email, u.Password, u.ID).
					WillReturnResult(sqlmock.NewErrorResult(testErr))
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepare(tc.args)

			err := UpdateUser(db, tc.args)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error to be %v, but got %v", tc.expectedError, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
