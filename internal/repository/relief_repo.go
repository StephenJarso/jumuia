package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateRelief(db *sql.DB, relief models.Relief) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO relief(member_id, disaster_id, type, amount, date_given) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		relief.MemberID,
		relief.DisasterID,
		relief.Type,
		relief.Amount,
		relief.DateGiven,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllRelief(db *sql.DB) ([]models.Relief, error) {
	rows, err := db.Query("SELECT id, member_id, disaster_id, type, amount, date_given FROM relief")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var relief []models.Relief
	for rows.Next() {
		var r models.Relief
		err := rows.Scan(&r.ID, &r.MemberID, &r.DisasterID, &r.Type, &r.Amount, &r.DateGiven)
		if err != nil {
			return nil, err
		}
		relief = append(relief, r)
	}
	return relief, nil
}

func GetReliefByMember(db *sql.DB, memberID int) ([]models.Relief, error) {
	rows, err := db.Query("SELECT id, member_id, disaster_id, type, amount, date_given FROM relief WHERE member_id = ?", memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var relief []models.Relief
	for rows.Next() {
		var r models.Relief
		err := rows.Scan(&r.ID, &r.MemberID, &r.DisasterID, &r.Type, &r.Amount, &r.DateGiven)
		if err != nil {
			return nil, err
		}
		relief = append(relief, r)
	}
	return relief, nil
}
