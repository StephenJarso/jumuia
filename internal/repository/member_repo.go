package repository

import (
	"database/sql"
	"jumuia/internal/models"
)

func AddMembers(db *sql.DB,member models.Member)(int64,error){
stmt,err:=db.Prepare("INSERT INTO members(group_id,name,phone,role) VALUES(?, ?, ?, ?)")
if err != nil{
	return 0, err
}
defer stmt.Close()
res,err:=stmt.Exec(
	member.GroupId,
	member.Name,
	member.Phone,
	member.Role,
)
if err!=nil{
	return 0, err
}
return res.LastInsertId()
}