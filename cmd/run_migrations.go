package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"post-comment-service/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %v", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	migrationsPath := filepath.Join(currentDir, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath)),
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	logger.Info.Println("Migrations applied successfully")
	return nil
}
