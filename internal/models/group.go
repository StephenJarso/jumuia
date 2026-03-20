package models

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Village  string `json:"village"`
	District string `json:"district"`
	// CreatedAt string `json:"createdAt"`
}
