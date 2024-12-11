package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	dropTableSQL := `DROP TABLE IF EXISTS users;`
	_, err = DB.Exec(dropTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	createTableSQL := `
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	log.Println("Database initialized.")
}
