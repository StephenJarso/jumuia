package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new savings entry
func NewSavingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := repository.GetAllMembers(db)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}
		data := struct {
			Members   []models.Member
			CSRFToken string
		}{
			Members:   members,
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/savings_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save savings
func CreateSavingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/savings/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		memberIDStr := r.FormValue("member_id")
		amountStr := r.FormValue("amount")
		meetingDate := r.FormValue("meeting_date")

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

		meetingDate, ok = middleware.ValidateRequired(meetingDate, "meeting_date")
		if !ok {
			middleware.WriteValidationError(w, "Meeting date is required")
			return
		}

		savings := models.Savings{
			MemberID:    memberID,
			Amount:      amount,
			MeetingDate: meetingDate,
		}
		_, err := repository.CreateSavings(db, savings)
		if err != nil {
			http.Error(w, "Error saving savings", 500)
			return
		}
		http.Redirect(w, r, "/savings", http.StatusSeeOther)
	}
}

// List all savings
func ListSavingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		savings, err := repository.GetAllSavings(db)
		if err != nil {
			http.Error(w, "Error fetching savings", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/savings_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, savings)
	}
}
