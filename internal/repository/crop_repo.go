package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateCrop(db *sql.DB, crop models.Crop) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO crops(name) VALUES(?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(crop.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllCrops(db *sql.DB) ([]models.Crop, error) {
	rows, err := db.Query("SELECT id, name FROM crops")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var crops []models.Crop
	for rows.Next() {
		var c models.Crop
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		crops = append(crops, c)
	}
	return crops, nil
}
