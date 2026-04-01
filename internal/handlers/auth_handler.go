package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// Show login form
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Validate required fields
			username, ok := middleware.ValidateRequired(username, "username")
			if !ok {
				middleware.WriteValidationError(w, "Username is required")
				return
			}

			password, ok = middleware.ValidateRequired(password, "password")
			if !ok {
				middleware.WriteValidationError(w, "Password is required")
				return
			}

			// Get user from database
			user, err := repository.GetUserByUsername(db, username)
			if err != nil {
				middleware.WriteValidationError(w, "Invalid username or password")
				return
			}

			// Check password
			err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
			if err != nil {
				middleware.WriteValidationError(w, "Invalid username or password")
				return
			}

			// Set session cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "user_id",
				Value:    string(rune(user.ID)),
				Path:     "/",
				MaxAge:   86400 * 7, // 7 days
				HttpOnly: true,
				Secure:   false, // Set to true in production with HTTPS
			})

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}

		// Show login form
		tmpl, err := template.New("login.html").Funcs(funcMap).ParseFiles("web/templates/login.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}
		tmpl.Execute(w, nil)
	}
}

// Show register form
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")
			confirmPassword := r.FormValue("confirm_password")
			email := r.FormValue("email")

			// Validate required fields
			username, ok := middleware.ValidateRequired(username, "username")
			if !ok {
				middleware.WriteValidationError(w, "Username is required")
				return
			}

			password, ok = middleware.ValidateRequired(password, "password")
			if !ok {
				middleware.WriteValidationError(w, "Password is required")
				return
			}

			if password != confirmPassword {
				middleware.WriteValidationError(w, "Passwords do not match")
				return
			}

			if len(password) < 6 {
				middleware.WriteValidationError(w, "Password must be at least 6 characters")
				return
			}

			// Hash password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Error hashing password", 500)
				return
			}

			// Create user
			user := models.User{
				Username:     username,
				PasswordHash: string(hashedPassword),
				Email:        email,
				Role:         "user",
			}

			_, err = repository.CreateUser(db, user)
			if err != nil {
				middleware.WriteValidationError(w, "Username already exists")
				return
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Show register form
		tmpl, err := template.New("register.html").Funcs(funcMap).ParseFiles("web/templates/register.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}
		tmpl.Execute(w, nil)
	}
}

// Logout handler
func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clear session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "user_id",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
