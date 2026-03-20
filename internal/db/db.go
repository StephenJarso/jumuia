package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB opens the sqlite database and returns a *sql.DB
func InitDB() *sql.DB {
	//open database(create file if it does not exist)
	db, err := sql.Open("sqlite3", "./jumuia.db")
	if err != nil {
		log.Fatal("Failed to open database", err)
	}
	//Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database", err)
	}
	// Enable foreign keys in Sqlite (important)
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}
	log.Println("Database connected and foreign keys enabled")
	return db
}
