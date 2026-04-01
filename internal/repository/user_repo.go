package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user models.User) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(username, password_hash, email, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	resp, err := stmt.Exec(user.Username, user.PasswordHash, user.Email, user.Role)
	if err != nil {
		return 0, err
	}
	return resp.LastInsertId()
}

// GetUserByUsername fetches a user by username
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, email, role, created_at
		FROM users
		WHERE username = ?
	`
	var u models.User
	err := db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByID fetches a user by ID
func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	query := `
		SELECT id, username, password_hash, email, role, created_at
		FROM users
		WHERE id = ?
	`
	var u models.User
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetAllUsers fetches all users from the database
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	query := `
		SELECT id, username, password_hash, email, role, created_at
		FROM users
		ORDER BY username
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Email, &u.Role, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
