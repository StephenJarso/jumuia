package models

type Member struct{
	ID int `json:"id"`
	GroupId int `json:"group_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	JoinedAt string `json:"joined_at"`
}