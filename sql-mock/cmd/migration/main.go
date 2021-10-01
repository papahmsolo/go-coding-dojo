package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("cannot start app: %v", err)
	}
	log.Print("migration ran successfully")
}

func run() error {
	// db, err := sql.Open("postgres", "")
	// if err != nil {
	// 	return fmt.Errorf("cannot oped db: %w", err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	return fmt.Errorf("cannot ping db: %w", err)
	// }

	m, err := migrate.New(
		"file://migrations",
		"postgres://guest:guest@localhost:2345/test_db?sslmode=disable")

	if err != nil {
		return fmt.Errorf("cannot create migration instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("cannot migrate schema: %w", err)
	}

	// err = m.Steps(-1)
	// if err != nil {
	// 	return fmt.Errorf("cannot downgrade schema by 1: %w", err)
	// }

	return nil
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r Repo) InTx(f func(*sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %w", err)
	}

	err = f(tx)
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			return fmt.Errorf("cannot rollback transaction: %v, original err: %w", rerr, err)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("cannot commit transaction: %w", err)
	}
	return nil
}

func (r Repo) CreateFile(tx *sql.Tx, fileName string) (string, error) {
	const q1 = "INSERT name INTO file VALUES ($1) RETURNING id"

	var fileID string
	err := tx.QueryRow(q1, fileName).Scan(&fileID)
	if err != nil {
		return "", fmt.Errorf("cannot insert file: %w", err)
	}
	return fileID, nil
}

func (r Repo) CreateRequest(tx *sql.Tx, fileID string) error {
	const q1 = "INSERT file_id INTO request VALUES ($1)"

	_, err := tx.Exec(q1, fileID)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}

	return nil
}

func (r Repo) CreateRequestInTx(fileName string) error {
	err := r.InTx(func(tx *sql.Tx) error {
		fileID, err := r.CreateFile(tx, fileName)
		if err != nil {
			return fmt.Errorf("cannot create file in transaction: %w", err)
		}

		err = r.CreateRequest(tx, fileID)
		if err != nil {
			return fmt.Errorf("cannot create request in transaction: %w", err)
		}
		return nil
	})
	return err
}
