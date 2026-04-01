package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new repayment
func NewRepaymentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loans, err := repository.GetAllLoans(db)
		if err != nil {
			http.Error(w, "Error loading loans", 500)
			return
		}
		data := struct {
			Loans     []models.Loan
			CSRFToken string
		}{
			Loans:     loans,
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/repayment_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save repayment
func CreateRepaymentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/repayments/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		loanIDStr := r.FormValue("loan_id")
		amountStr := r.FormValue("amount")
		paymentDate := r.FormValue("payment_date")

		loanID, ok := middleware.ValidateInt(loanIDStr, "loan_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid loan ID")
			return
		}

		amount, ok := middleware.ValidateFloat(amountStr, "amount")
		if !ok || amount <= 0 {
			middleware.WriteValidationError(w, "Invalid amount")
			return
		}

		paymentDate, ok = middleware.ValidateRequired(paymentDate, "payment_date")
		if !ok {
			middleware.WriteValidationError(w, "Payment date is required")
			return
		}

		repayment := models.Repayment{
			LoanID:      loanID,
			Amount:      amount,
			PaymentDate: paymentDate,
		}
		_, err := repository.CreateRepayment(db, repayment)
		if err != nil {
			http.Error(w, "Error saving repayment", 500)
			return
		}
		http.Redirect(w, r, "/repayments", http.StatusSeeOther)
	}
}

// List all repayments
func ListRepaymentsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repayments, err := repository.GetAllRepayments(db)
		if err != nil {
			http.Error(w, "Error fetching repayments", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/repayments_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, repayments)
	}
}
