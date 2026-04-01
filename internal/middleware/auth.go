package middleware

import (
	"net/http"
)

// RequireAuth middleware checks if user is authenticated
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for user_id cookie
		cookie, err := r.Cookie("user_id")
		if err != nil || cookie.Value == "" {
			// Redirect to login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// User is authenticated, proceed with the request
		next(w, r)
	}
}
