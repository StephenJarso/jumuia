package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new loan
func NewLoanHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := repository.GetAllMembers(db)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}
		seasons, err := repository.GetAllSeasons(db)
		if err != nil {
			http.Error(w, "Error loading seasons", 500)
			return
		}
		data := struct {
			Members   []models.Member
			Seasons   []models.Season
			CSRFToken string
		}{
			Members:   members,
			Seasons:   seasons,
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/loan_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save loan
func CreateLoanHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/loans/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		memberIDStr := r.FormValue("member_id")
		amountStr := r.FormValue("amount")
		purpose := r.FormValue("purpose")
		status := r.FormValue("status")
		issuedDate := r.FormValue("issued_date")
		dueDate := r.FormValue("due_date")
		seasonIDStr := r.FormValue("season_id")

		memberID, ok := middleware.ValidateInt(memberIDStr, "member_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid member ID")
			return
		}

		amount, ok := middleware.ValidateFloat(amountStr, "amount")
		if !ok || amount <= 0 {
			middleware.WriteValidationError(w, "Invalid amount")
			return
		}

		purpose, ok = middleware.ValidateRequired(purpose, "purpose")
		if !ok {
			middleware.WriteValidationError(w, "Purpose is required")
			return
		}

		status, ok = middleware.ValidateRequired(status, "status")
		if !ok {
			middleware.WriteValidationError(w, "Status is required")
			return
		}

		issuedDate, ok = middleware.ValidateRequired(issuedDate, "issued_date")
		if !ok {
			middleware.WriteValidationError(w, "Issued date is required")
			return
		}

		dueDate, ok = middleware.ValidateRequired(dueDate, "due_date")
		if !ok {
			middleware.WriteValidationError(w, "Due date is required")
			return
		}

		seasonID, ok := middleware.ValidateInt(seasonIDStr, "season_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid season ID")
			return
		}

		loan := models.Loan{
			MemberID:   memberID,
			Amount:     amount,
			Purpose:    purpose,
			Status:     status,
			IssuedDate: issuedDate,
			DueDate:    dueDate,
			SeasonID:   seasonID,
		}
		_, err := repository.CreateLoan(db, loan)
		if err != nil {
			http.Error(w, "Error saving loan", 500)
			return
		}
		http.Redirect(w, r, "/loans", http.StatusSeeOther)
	}
}

// List all loans
func ListLoansHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loans, err := repository.GetAllLoans(db)
		if err != nil {
			http.Error(w, "Error fetching loans", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/loans_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, loans)
	}
}
