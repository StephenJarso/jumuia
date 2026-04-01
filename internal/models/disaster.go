package models

type Disaster struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
}
