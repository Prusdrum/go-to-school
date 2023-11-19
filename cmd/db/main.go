package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

func main() {
	// db, err := sql.Open("sqlite3", filepath.Join("..", "..", "database", "go-to-school.db"))
	db, err := sql.Open("sqlite", "file:../../database/go-to-school.db")
	if err != nil {
		log.Fatal("open db: ", err)
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatal("setup driver: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"sqlite", driver)
	if err != nil {
		log.Fatal("migrate db: ", err)
	}
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
