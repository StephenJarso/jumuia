package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateLoan(db *sql.DB, loan models.Loan) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO loans(member_id, amount, purpose, status, issued_date, due_date, season_id) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		loan.MemberID,
		loan.Amount,
		loan.Purpose,
		loan.Status,
		loan.IssuedDate,
		loan.DueDate,
		loan.SeasonID,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllLoans(db *sql.DB) ([]models.Loan, error) {
	rows, err := db.Query("SELECT id, member_id, amount, purpose, status, issued_date, due_date, season_id FROM loans")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var loans []models.Loan
	for rows.Next() {
		var l models.Loan
		err := rows.Scan(&l.ID, &l.MemberID, &l.Amount, &l.Purpose, &l.Status, &l.IssuedDate, &l.DueDate, &l.SeasonID)
		if err != nil {
			return nil, err
		}
		loans = append(loans, l)
	}
	return loans, nil
}

func GetLoansByMember(db *sql.DB, memberID int) ([]models.Loan, error) {
	rows, err := db.Query("SELECT id, member_id, amount, purpose, status, issued_date, due_date, season_id FROM loans WHERE member_id = ?", memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var loans []models.Loan
	for rows.Next() {
		var l models.Loan
		err := rows.Scan(&l.ID, &l.MemberID, &l.Amount, &l.Purpose, &l.Status, &l.IssuedDate, &l.DueDate, &l.SeasonID)
		if err != nil {
			return nil, err
		}
		loans = append(loans, l)
	}
	return loans, nil
}
