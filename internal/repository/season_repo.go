package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateSeason(db *sql.DB, season models.Season) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO seasons(name, year) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(season.Name, season.Year)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllSeasons(db *sql.DB) ([]models.Season, error) {
	rows, err := db.Query("SELECT id, name, year FROM seasons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var seasons []models.Season
	for rows.Next() {
		var s models.Season
		err := rows.Scan(&s.ID, &s.Name, &s.Year)
		if err != nil {
			return nil, err
		}
		seasons = append(seasons, s)
	}
	return seasons, nil
}
