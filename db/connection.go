package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(fmt.Sprintf("Could not create event table: %s", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    location VARCHAR(255) NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL
)`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event table.")
	}
}
