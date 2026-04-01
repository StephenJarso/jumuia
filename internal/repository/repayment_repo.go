package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateRepayment(db *sql.DB, repayment models.Repayment) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO repayments(loan_id, amount, payment_date) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		repayment.LoanID,
		repayment.Amount,
		repayment.PaymentDate,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetAllRepayments(db *sql.DB) ([]models.Repayment, error) {
	rows, err := db.Query("SELECT id, loan_id, amount, payment_date FROM repayments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repayments []models.Repayment
	for rows.Next() {
		var r models.Repayment
		err := rows.Scan(&r.ID, &r.LoanID, &r.Amount, &r.PaymentDate)
		if err != nil {
			return nil, err
		}
		repayments = append(repayments, r)
	}
	return repayments, nil
}

func GetRepaymentsByLoan(db *sql.DB, loanID int) ([]models.Repayment, error) {
	rows, err := db.Query("SELECT id, loan_id, amount, payment_date FROM repayments WHERE loan_id = ?", loanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repayments []models.Repayment
	for rows.Next() {
		var r models.Repayment
		err := rows.Scan(&r.ID, &r.LoanID, &r.Amount, &r.PaymentDate)
		if err != nil {
			return nil, err
		}
		repayments = append(repayments, r)
	}
	return repayments, nil
}
