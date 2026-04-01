package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

// CreateGroup inserts a new group into our database
func CreateGroup(db *sql.DB, group models.Group) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO groups(name, village, district, leader_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	resp, err := stmt.Exec(group.Name, group.Village, group.District, group.LeaderID)
	if err != nil {
		return 0, err
	}
	return resp.LastInsertId()
}

// GetAllGroups fetches all groups from the database with leader information
func GetAllGroups(db *sql.DB) ([]models.GroupWithLeader, error) {
	query := `
		SELECT g.id, g.name, g.village, g.district, g.leader_id, g.created_at,
			   COALESCE(m.name, 'No Leader') as leader_name,
			   COALESCE(m.phone, '') as leader_phone,
			   (SELECT COUNT(*) FROM members WHERE group_id = g.id) as member_count
		FROM groups g
		LEFT JOIN members m ON g.leader_id = m.id
		ORDER BY g.name
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groups []models.GroupWithLeader
	for rows.Next() {
		var g models.GroupWithLeader
		err := rows.Scan(&g.ID, &g.Name, &g.Village, &g.District, &g.LeaderID, &g.CreatedAt,
			&g.LeaderName, &g.LeaderPhone, &g.MemberCount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	return groups, nil
}

// GetGroupByID fetches a single group by ID with leader information
func GetGroupByID(db *sql.DB, id int) (*models.GroupWithLeader, error) {
	query := `
		SELECT g.id, g.name, g.village, g.district, g.leader_id, g.created_at,
			   COALESCE(m.name, 'No Leader') as leader_name,
			   COALESCE(m.phone, '') as leader_phone,
			   (SELECT COUNT(*) FROM members WHERE group_id = g.id) as member_count
		FROM groups g
		LEFT JOIN members m ON g.leader_id = m.id
		WHERE g.id = ?
	`
	var g models.GroupWithLeader
	err := db.QueryRow(query, id).Scan(&g.ID, &g.Name, &g.Village, &g.District, &g.LeaderID, &g.CreatedAt,
		&g.LeaderName, &g.LeaderPhone, &g.MemberCount)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// UpdateGroupLeader updates the leader of a group
func UpdateGroupLeader(db *sql.DB, groupID int, leaderID int) error {
	stmt, err := db.Prepare("UPDATE groups SET leader_id = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(leaderID, groupID)
	return err
}

// GetGroupStats returns statistics for a specific group
func GetGroupStats(db *sql.DB, groupID int) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Get total savings
	var totalSavings float64
	err := db.QueryRow(`
		SELECT COALESCE(SUM(s.amount), 0)
		FROM savings s
		JOIN members m ON s.member_id = m.id
		WHERE m.group_id = ?
	`, groupID).Scan(&totalSavings)
	if err != nil {
		return nil, err
	}
	stats["total_savings"] = totalSavings

	// Get total loans
	var totalLoans float64
	err = db.QueryRow(`
		SELECT COALESCE(SUM(l.amount), 0)
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ? AND l.status = 'active'
	`, groupID).Scan(&totalLoans)
	if err != nil {
		return nil, err
	}
	stats["total_loans"] = totalLoans

	// Get total repayments
	var totalRepayments float64
	err = db.QueryRow(`
		SELECT COALESCE(SUM(r.amount), 0)
		FROM repayments r
		JOIN loans l ON r.loan_id = l.id
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ?
	`, groupID).Scan(&totalRepayments)
	if err != nil {
		return nil, err
	}
	stats["total_repayments"] = totalRepayments

	// Get active loans count
	var activeLoans int
	err = db.QueryRow(`
		SELECT COUNT(*)
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ? AND l.status = 'active'
	`, groupID).Scan(&activeLoans)
	if err != nil {
		return nil, err
	}
	stats["active_loans"] = activeLoans

	// Get total relief
	var totalRelief float64
	err = db.QueryRow(`
		SELECT COALESCE(SUM(r.amount), 0)
		FROM relief r
		JOIN members m ON r.member_id = m.id
		WHERE m.group_id = ?
	`, groupID).Scan(&totalRelief)
	if err != nil {
		return nil, err
	}
	stats["total_relief"] = totalRelief

	// Get pending loans count
	var pendingLoans int
	err = db.QueryRow(`
		SELECT COUNT(*)
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ? AND l.status = 'pending'
	`, groupID).Scan(&pendingLoans)
	if err != nil {
		return nil, err
	}
	stats["pending_loans"] = pendingLoans

	// Get completed loans count
	var completedLoans int
	err = db.QueryRow(`
		SELECT COUNT(*)
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ? AND l.status = 'completed'
	`, groupID).Scan(&completedLoans)
	if err != nil {
		return nil, err
	}
	stats["completed_loans"] = completedLoans

	return stats, nil
}

// GetGroupSavings returns all savings for a specific group
func GetGroupSavings(db *sql.DB, groupID int) ([]map[string]interface{}, error) {
	query := `
		SELECT s.id, s.amount, s.meeting_date, m.name as member_name
		FROM savings s
		JOIN members m ON s.member_id = m.id
		WHERE m.group_id = ?
		ORDER BY s.meeting_date DESC
	`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var savings []map[string]interface{}
	for rows.Next() {
		var id int
		var amount float64
		var meetingDate, memberName string
		err := rows.Scan(&id, &amount, &meetingDate, &memberName)
		if err != nil {
			return nil, err
		}
		savings = append(savings, map[string]interface{}{
			"id":           id,
			"amount":       amount,
			"meeting_date": meetingDate,
			"member_name":  memberName,
		})
	}
	return savings, nil
}

// GetGroupRelief returns all relief for a specific group
func GetGroupRelief(db *sql.DB, groupID int) ([]map[string]interface{}, error) {
	query := `
		SELECT r.id, r.type, r.amount, r.date_given, m.name as member_name, d.description as disaster_description
		FROM relief r
		JOIN members m ON r.member_id = m.id
		LEFT JOIN disasters d ON r.disaster_id = d.id
		WHERE m.group_id = ?
		ORDER BY r.date_given DESC
	`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relief []map[string]interface{}
	for rows.Next() {
		var id int
		var reliefType, memberName string
		var amount float64
		var dateGiven, disasterDescription string
		err := rows.Scan(&id, &reliefType, &amount, &dateGiven, &memberName, &disasterDescription)
		if err != nil {
			return nil, err
		}
		relief = append(relief, map[string]interface{}{
			"id":                   id,
			"type":                 reliefType,
			"amount":               amount,
			"date_given":           dateGiven,
			"member_name":          memberName,
			"disaster_description": disasterDescription,
		})
	}
	return relief, nil
}

// GetGroupLoans returns all loans for a specific group
func GetGroupLoans(db *sql.DB, groupID int) ([]map[string]interface{}, error) {
	query := `
		SELECT l.id, l.amount, l.purpose, l.status, l.issued_date, l.due_date, m.name as member_name
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ?
		ORDER BY l.issued_date DESC
	`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loans []map[string]interface{}
	for rows.Next() {
		var id int
		var amount float64
		var purpose, status, memberName string
		var issuedDate, dueDate string
		err := rows.Scan(&id, &amount, &purpose, &status, &issuedDate, &dueDate, &memberName)
		if err != nil {
			return nil, err
		}
		loans = append(loans, map[string]interface{}{
			"id":          id,
			"amount":      amount,
			"purpose":     purpose,
			"status":      status,
			"issued_date": issuedDate,
			"due_date":    dueDate,
			"member_name": memberName,
		})
	}
	return loans, nil
}

// GetGroupSavingsByMember returns savings aggregated by member for a specific group
func GetGroupSavingsByMember(db *sql.DB, groupID int) ([]map[string]interface{}, error) {
	query := `
		SELECT m.name, COALESCE(SUM(s.amount), 0) as total_savings
		FROM members m
		LEFT JOIN savings s ON m.id = s.member_id
		WHERE m.group_id = ?
		GROUP BY m.id, m.name
		ORDER BY total_savings DESC
	`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var savingsByMember []map[string]interface{}
	for rows.Next() {
		var name string
		var totalSavings float64
		err := rows.Scan(&name, &totalSavings)
		if err != nil {
			return nil, err
		}
		savingsByMember = append(savingsByMember, map[string]interface{}{
			"member_name":   name,
			"total_savings": totalSavings,
		})
	}
	return savingsByMember, nil
}

// GetGroupLoansByStatus returns loans aggregated by status for a specific group
func GetGroupLoansByStatus(db *sql.DB, groupID int) ([]map[string]interface{}, error) {
	query := `
		SELECT l.status, COUNT(*) as count, COALESCE(SUM(l.amount), 0) as total_amount
		FROM loans l
		JOIN members m ON l.member_id = m.id
		WHERE m.group_id = ?
		GROUP BY l.status
	`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loansByStatus []map[string]interface{}
	for rows.Next() {
		var status string
		var count int
		var totalAmount float64
		err := rows.Scan(&status, &count, &totalAmount)
		if err != nil {
			return nil, err
		}
		loansByStatus = append(loansByStatus, map[string]interface{}{
			"status":       status,
			"count":        count,
			"total_amount": totalAmount,
		})
	}
	return loansByStatus, nil
}
