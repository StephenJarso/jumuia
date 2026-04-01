package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateSavings(db *sql.DB, savings models.Savings) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO savings(member_id, amount, meeting_date) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		savings.MemberID,
		savings.Amount,
		savings.MeetingDate,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllSavings(db *sql.DB) ([]models.Savings, error) {
	rows, err := db.Query("SELECT id, member_id, amount, meeting_date FROM savings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var savings []models.Savings
	for rows.Next() {
		var s models.Savings
		err := rows.Scan(&s.ID, &s.MemberID, &s.Amount, &s.MeetingDate)
		if err != nil {
			return nil, err
		}
		savings = append(savings, s)
	}
	return savings, nil
}

func GetSavingsByMember(db *sql.DB, memberID int) ([]models.Savings, error) {
	rows, err := db.Query("SELECT id, member_id, amount, meeting_date FROM savings WHERE member_id = ?", memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var savings []models.Savings
	for rows.Next() {
		var s models.Savings
		err := rows.Scan(&s.ID, &s.MemberID, &s.Amount, &s.MeetingDate)
		if err != nil {
			return nil, err
		}
		savings = append(savings, s)
	}
	return savings, nil
}
