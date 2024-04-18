package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationPath, migrationTable string
	flag.StringVar(&storagePath, "storage-path", "./storage/storage.db", "Path to the storage directory")
	flag.StringVar(&migrationPath, "migration-path", "./migration", "Path to the migration directory")
	flag.StringVar(&migrationTable, "migration-table", "migration", "Table to migrate")


	m, err := migrate.New(
		"file://"+ migrationPath, 
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationTable),
	)
	 
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations were applied to the database")
			return
		}
		panic(err)
	}

	fmt.Println("Migrations applied to the database successfully")
}