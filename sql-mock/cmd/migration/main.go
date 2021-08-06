package main

import (
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
