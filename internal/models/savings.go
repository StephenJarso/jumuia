package models

type Savings struct {
	ID          int     `json:"id"`
	MemberID    int     `json:"member_id"`
	Amount      float64 `json:"amount"`
	MeetingDate string  `json:"meeting_date"`
}
