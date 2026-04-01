package models

type Repayment struct {
	ID          int     `json:"id"`
	LoanID      int     `json:"loan_id"`
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
}
