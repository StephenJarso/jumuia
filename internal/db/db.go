package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB opens the sqlite database and returns a *sql.DB
func InitDB() *sql.DB {
	// Get the directory where the executable is located
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory", err)
	}

	// Create database path
	dbPath := filepath.Join(execDir, "jumuia.db")

	// Check if database exists and has data
	_, err = os.Stat(dbPath)
	if err == nil {
		// Database exists, check if it has data
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal("Failed to open database", err)
		}
		defer db.Close()

		var groupCount int
		err = db.QueryRow("SELECT COUNT(*) FROM groups").Scan(&groupCount)
		if err == nil && groupCount > 0 {
			log.Printf("Database already exists with %d groups", groupCount)
			// Reopen and return
			db, err = sql.Open("sqlite3", dbPath)
			if err != nil {
				log.Fatal("Failed to open database", err)
			}
			// Enable foreign keys
			_, err = db.Exec("PRAGMA foreign_keys = ON;")
			if err != nil {
				log.Fatal("Failed to enable foreign keys:", err)
			}
			return db
		}
	}

	// Open database (create file if it does not exist)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database", err)
	}
	// Test connection
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

	// Load schema and mock data
	log.Println("Loading schema and mock data...")
	loadSchemaAndMockData(db, execDir)

	return db
}

// loadSchemaAndMockData loads the schema and mock data into the database
func loadSchemaAndMockData(db *sql.DB, basePath string) {
	// Read schema file
	schemaPath := filepath.Join(basePath, "migrations", "schema.sql")
	schemaSQL, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Printf("Warning: Could not read schema file from %s: %v", schemaPath, err)
		return
	}

	// Execute schema
	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		log.Printf("Warning: Could not execute schema: %v", err)
		return
	}
	log.Println("Schema loaded successfully")

	// Read mock data file
	mockPath := filepath.Join(basePath, "migrations", "mock_data.sql")
	mockSQL, err := os.ReadFile(mockPath)
	if err != nil {
		log.Printf("Warning: Could not read mock data file from %s: %v", mockPath, err)
		return
	}

	// Execute mock data
	_, err = db.Exec(string(mockSQL))
	if err != nil {
		log.Printf("Warning: Could not execute mock data: %v", err)
		return
	}
	log.Println("Mock data loaded successfully")

	// Verify data was loaded
	var groupCount int
	err = db.QueryRow("SELECT COUNT(*) FROM groups").Scan(&groupCount)
	if err == nil {
		log.Printf("Successfully loaded %d groups", groupCount)
	}
}
