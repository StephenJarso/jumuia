package models

type Relief struct {
	ID         int     `json:"id"`
	MemberID   int     `json:"member_id"`
	DisasterID int     `json:"disaster_id"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
	DateGiven  string  `json:"date_given"`
}
