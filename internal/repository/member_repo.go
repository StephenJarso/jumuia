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

// GetMembers by group returns members of a group
func GetAllMembers(db *sql.DB) ([]models.Member, error) {
	rows, err := db.Query("SELECT id, group_id, name, phone, role FROM members")
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
		)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

// GetMembers by group returns members of a group
func GetMembersByGroup(db *sql.DB, groupID int) ([]models.Member, error) {
	rows, err := db.Query(`SELECT id,group_id,name,phone,role 
	FROM members 
	WHERE group_id = ?`, groupID)
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
		)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}
