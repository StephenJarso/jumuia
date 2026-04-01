package models

type Loan struct {
	ID         int     `json:"id"`
	MemberID   int     `json:"member_id"`
	Amount     float64 `json:"amount"`
	Purpose    string  `json:"purpose"`
	Status     string  `json:"status"`
	IssuedDate string  `json:"issued_date"`
	DueDate    string  `json:"due_date"`
	SeasonID   int     `json:"season_id"`
}
