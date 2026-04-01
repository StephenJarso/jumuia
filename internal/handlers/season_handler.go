package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new season
func NewSeasonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			CSRFToken string
		}{
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/season_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save season
func CreateSeasonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/seasons/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		name := r.FormValue("name")
		year := r.FormValue("year")

		name, ok := middleware.ValidateRequired(name, "name")
		if !ok {
			middleware.WriteValidationError(w, "Name is required")
			return
		}

		yearInt, ok := middleware.ValidateInt(year, "year")
		if !ok {
			middleware.WriteValidationError(w, "Invalid year")
			return
		}

		season := models.Season{
			Name: name,
			Year: yearInt,
		}
		_, err := repository.CreateSeason(db, season)
		if err != nil {
			http.Error(w, "Error saving season", 500)
			return
		}
		http.Redirect(w, r, "/seasons", http.StatusSeeOther)
	}
}

// List all seasons
func ListSeasonsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		seasons, err := repository.GetAllSeasons(db)
		if err != nil {
			http.Error(w, "Error fetching seasons", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/seasons_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, seasons)
	}
}
