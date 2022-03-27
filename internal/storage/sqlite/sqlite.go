package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func New(dbFilePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

func Migrate(db *sql.DB, migrationFilesPath string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
		NoTxWrap: true,
	})
	if err != nil {
		return fmt.Errorf("failed to init migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationFilesPath, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("failed to init migration instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
