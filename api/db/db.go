package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func Init() {

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		fmt.Println(err)
		panic("Could establish DB connection.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		userId INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create tables.")
	}
}
