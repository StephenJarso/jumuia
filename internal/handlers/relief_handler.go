package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new relief entry
func NewReliefHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := repository.GetAllMembers(db)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}
		disasters, err := repository.GetAllDisasters(db)
		if err != nil {
			http.Error(w, "Error loading disasters", 500)
			return
		}
		data := struct {
			Members   []models.Member
			Disasters []models.Disaster
			CSRFToken string
		}{
			Members:   members,
			Disasters: disasters,
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/relief_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save relief
func CreateReliefHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/relief/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		memberIDStr := r.FormValue("member_id")
		disasterIDStr := r.FormValue("disaster_id")
		reliefType := r.FormValue("type")
		amountStr := r.FormValue("amount")
		dateGiven := r.FormValue("date_given")

		memberID, ok := middleware.ValidateInt(memberIDStr, "member_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid member ID")
			return
		}

		disasterID, ok := middleware.ValidateInt(disasterIDStr, "disaster_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid disaster ID")
			return
		}

		reliefType, ok = middleware.ValidateRequired(reliefType, "type")
		if !ok {
			middleware.WriteValidationError(w, "Type is required")
			return
		}

		amount, ok := middleware.ValidateFloat(amountStr, "amount")
		if !ok || amount <= 0 {
			middleware.WriteValidationError(w, "Invalid amount")
			return
		}

		dateGiven, ok = middleware.ValidateRequired(dateGiven, "date_given")
		if !ok {
			middleware.WriteValidationError(w, "Date given is required")
			return
		}

		relief := models.Relief{
			MemberID:   memberID,
			DisasterID: disasterID,
			Type:       reliefType,
			Amount:     amount,
			DateGiven:  dateGiven,
		}
		_, err := repository.CreateRelief(db, relief)
		if err != nil {
			http.Error(w, "Error saving relief", 500)
			return
		}
		http.Redirect(w, r, "/relief", http.StatusSeeOther)
	}
}

// List all relief entries
func ListReliefHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		relief, err := repository.GetAllRelief(db)
		if err != nil {
			http.Error(w, "Error fetching relief", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/relief_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, relief)
	}
}
