package main

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path, err := filepath.Abs("database/go-to-school.sqlite3")
	if err != nil {
		log.Fatal("get db path: ", err)
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("open db: ", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal("setup driver: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration",
		"sqlite3", driver)

	if err != nil {
		log.Fatal("migrate db: ", err)
	}
	m.Up()
}
