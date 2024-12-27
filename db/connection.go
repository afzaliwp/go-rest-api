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
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user table.")
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255),
        location VARCHAR(255) NOT NULL,
        date_time DATETIME NOT NULL,
        user_id INTEGER NOT NULL,
        created_at DATETIME NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id))`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event table: " + err.Error())
	}
}
