package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new disaster
func NewDisasterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			CSRFToken string
		}{
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/disaster_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save disaster
func CreateDisasterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/disasters/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		disasterType := r.FormValue("type")
		description := r.FormValue("description")
		location := r.FormValue("location")
		date := r.FormValue("date")

		disasterType, ok := middleware.ValidateRequired(disasterType, "type")
		if !ok {
			middleware.WriteValidationError(w, "Type is required")
			return
		}

		description, ok = middleware.ValidateRequired(description, "description")
		if !ok {
			middleware.WriteValidationError(w, "Description is required")
			return
		}

		location, ok = middleware.ValidateRequired(location, "location")
		if !ok {
			middleware.WriteValidationError(w, "Location is required")
			return
		}

		date, ok = middleware.ValidateRequired(date, "date")
		if !ok {
			middleware.WriteValidationError(w, "Date is required")
			return
		}

		disaster := models.Disaster{
			Type:        disasterType,
			Description: description,
			Location:    location,
			Date:        date,
		}
		_, err := repository.CreateDisaster(db, disaster)
		if err != nil {
			http.Error(w, "Error saving disaster", 500)
			return
		}
		http.Redirect(w, r, "/disasters", http.StatusSeeOther)
	}
}

// List all disasters
func ListDisastersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		disasters, err := repository.GetAllDisasters(db)
		if err != nil {
			http.Error(w, "Error fetching disasters", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/disasters_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, disasters)
	}
}
