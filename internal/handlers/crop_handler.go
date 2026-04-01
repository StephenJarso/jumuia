package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show the form to create a new crop
func NewCropHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			CSRFToken string
		}{
			CSRFToken: middleware.GenerateCSRFToken(),
		}
		tmpl, err := template.ParseFiles("web/templates/crop_form.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save crop
func CreateCropHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/crops/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		name := r.FormValue("name")
		name, ok := middleware.ValidateRequired(name, "name")
		if !ok {
			middleware.WriteValidationError(w, "Name is required")
			return
		}

		crop := models.Crop{
			Name: name,
		}
		_, err := repository.CreateCrop(db, crop)
		if err != nil {
			http.Error(w, "Error saving crop", 500)
			return
		}
		http.Redirect(w, r, "/crops", http.StatusSeeOther)
	}
}

// List all crops
func ListCropsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		crops, err := repository.GetAllCrops(db)
		if err != nil {
			http.Error(w, "Error fetching crops", 500)
			return
		}

		tmpl, err := template.ParseFiles("web/templates/crops_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, crops)
	}
}
