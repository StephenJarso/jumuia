package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateDisaster(db *sql.DB, disaster models.Disaster) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO disasters(type, description, location, date) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		disaster.Type,
		disaster.Description,
		disaster.Location,
		disaster.Date,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllDisasters(db *sql.DB) ([]models.Disaster, error) {
	rows, err := db.Query("SELECT id, type, description, location, date FROM disasters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var disasters []models.Disaster
	for rows.Next() {
		var d models.Disaster
		err := rows.Scan(&d.ID, &d.Type, &d.Description, &d.Location, &d.Date)
		if err != nil {
			return nil, err
		}
		disasters = append(disasters, d)
	}
	return disasters, nil
}
