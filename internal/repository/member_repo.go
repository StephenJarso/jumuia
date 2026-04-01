package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

func CreateMember(db *sql.DB, member models.Member) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO members(group_id,name,phone,role) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		member.GroupId,
		member.Name,
		member.Phone,
		member.Role,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetAllMembers returns all members
func GetAllMembers(db *sql.DB) ([]models.Member, error) {
	rows, err := db.Query("SELECT id, group_id, name, phone, role, joined_at FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []models.Member
	for rows.Next() {
		var m models.Member
		err := rows.Scan(
			&m.ID,
			&m.GroupId,
			&m.Name,
			&m.Phone,
			&m.Role,
			&m.JoinedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

// GetMembersByGroup returns members of a group
func GetMembersByGroup(db *sql.DB, groupID int) ([]models.Member, error) {
	rows, err := db.Query(`SELECT id, group_id, name, phone, role, joined_at 
	FROM members 
	WHERE group_id = ?
	ORDER BY CASE role WHEN 'leader' THEN 1 WHEN 'treasurer' THEN 2 ELSE 3 END, name`, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []models.Member
	for rows.Next() {
		var m models.Member
		err := rows.Scan(
			&m.ID,
			&m.GroupId,
			&m.Name,
			&m.Phone,
			&m.Role,
			&m.JoinedAt,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

// GetMemberByID returns a single member by ID
func GetMemberByID(db *sql.DB, memberID int) (*models.Member, error) {
	var m models.Member
	err := db.QueryRow(`SELECT id, group_id, name, phone, role, joined_at 
	FROM members WHERE id = ?`, memberID).Scan(
		&m.ID,
		&m.GroupId,
		&m.Name,
		&m.Phone,
		&m.Role,
		&m.JoinedAt,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMemberFinancialSummary returns financial summary for a member
func GetMemberFinancialSummary(db *sql.DB, memberID int) (map[string]interface{}, error) {
	summary := make(map[string]interface{})

	// Get total savings
	var totalSavings float64
	err := db.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM savings WHERE member_id = ?", memberID).Scan(&totalSavings)
	if err != nil {
		return nil, err
	}
	summary["total_savings"] = totalSavings

	// Get total loans
	var totalLoans float64
	err = db.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM loans WHERE member_id = ? AND status = 'active'", memberID).Scan(&totalLoans)
	if err != nil {
		return nil, err
	}
	summary["total_loans"] = totalLoans

	// Get total repayments
	var totalRepayments float64
	err = db.QueryRow(`
		SELECT COALESCE(SUM(r.amount), 0)
		FROM repayments r
		JOIN loans l ON r.loan_id = l.id
		WHERE l.member_id = ?
	`, memberID).Scan(&totalRepayments)
	if err != nil {
		return nil, err
	}
	summary["total_repayments"] = totalRepayments

	// Get outstanding loan balance
	summary["outstanding_balance"] = totalLoans - totalRepayments

	// Get savings history
	savingsRows, err := db.Query("SELECT amount, meeting_date FROM savings WHERE member_id = ? ORDER BY meeting_date DESC LIMIT 10", memberID)
	if err != nil {
		return nil, err
	}
	defer savingsRows.Close()
	var savingsHistory []map[string]interface{}
	for savingsRows.Next() {
		var amount float64
		var date string
		err := savingsRows.Scan(&amount, &date)
		if err != nil {
			return nil, err
		}
		savingsHistory = append(savingsHistory, map[string]interface{}{
			"amount": amount,
			"date":   date,
		})
	}
	summary["savings_history"] = savingsHistory

	// Get active loans
	loanRows, err := db.Query("SELECT id, amount, purpose, issued_date, due_date FROM loans WHERE member_id = ? AND status = 'active' ORDER BY issued_date DESC", memberID)
	if err != nil {
		return nil, err
	}
	defer loanRows.Close()
	var activeLoans []map[string]interface{}
	for loanRows.Next() {
		var id int
		var amount float64
		var purpose, issuedDate, dueDate string
		err := loanRows.Scan(&id, &amount, &purpose, &issuedDate, &dueDate)
		if err != nil {
			return nil, err
		}
		activeLoans = append(activeLoans, map[string]interface{}{
			"id":          id,
			"amount":      amount,
			"purpose":     purpose,
			"issued_date": issuedDate,
			"due_date":    dueDate,
		})
	}
	summary["active_loans"] = activeLoans

	return summary, nil
}
