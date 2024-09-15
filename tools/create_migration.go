package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	name := flag.String("name", "", "Name of the migration")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please provide a name for the migration using the -name flag")
		os.Exit(1)
	}

	timestamp := time.Now().Unix()
	migrationsDir := "./migrations"

	err := os.MkdirAll(migrationsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create migrations directory: %v\n", err)
		os.Exit(1)
	}

	upFilename := fmt.Sprintf("%d_%s.up.sql", timestamp, *name)
	downFilename := fmt.Sprintf("%d_%s.down.sql", timestamp, *name)

	err = createFile(filepath.Join(migrationsDir, upFilename))
	if err != nil {
		fmt.Printf("Failed to create up migration file: %v\n", err)
		os.Exit(1)
	}

	err = createFile(filepath.Join(migrationsDir, downFilename))
	if err != nil {
		fmt.Printf("Failed to create down migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created migration files:\n%s\n%s\n", upFilename, downFilename)
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
