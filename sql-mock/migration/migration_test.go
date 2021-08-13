// +build integration

package migration

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMigration(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get path to migrations folder: %v", err)
	}

	m, err := migrate.New(
		"file://"+path.Dir(dir)+"/migrations",
		"postgres://guest:guest@localhost:2345/test_db?sslmode=disable")
	if err != nil {
		t.Fatalf("cannot create migration instance: %v", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		t.Fatalf("migrate up filed: %v", err)
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		t.Fatalf("migrate down filed: %v", err)
	}
}
