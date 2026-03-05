package repository

import (
	"database/sql"

	"jumuia/internal/models"
)

// createGroup inserts a new group into our database
func CreateGroup(db *sql.DB, group models.Group) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO groups(name,village,district) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	resp, err := stmt.Exec(group.Name, group.Village, group.District)
	if err != nil {
		return 0, err
	}
	// return id of the newly inserted row
	return resp.LastInsertId()
}

//GetAllGroups fetches all groups from the database

func GetAllGroups(db *sql.DB)([]models.Group,error){
	rows,err:=db.Query("SELECT id,name,village, district FROM groups")
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	var groups []models.Group
	for rows.Next(){
		var g models.Group
		err := rows.Scan(&g.ID,&g.Name,&g.Village,&g.District)
		if err != nil{
			return nil,err
		}
		groups = append(groups,g)
	}
return groups, nil
}
