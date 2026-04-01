package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

var (
	csrfTokens = make(map[string]time.Time)
	csrfMutex  sync.RWMutex
)

// GenerateCSRFToken creates a new CSRF token
func GenerateCSRFToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	token := hex.EncodeToString(bytes)

	csrfMutex.Lock()
	csrfTokens[token] = time.Now().Add(1 * time.Hour) // Token expires in 1 hour
	csrfMutex.Unlock()

	return token
}

// ValidateCSRFToken checks if a token is valid
func ValidateCSRFToken(token string) bool {
	csrfMutex.RLock()
	expiry, exists := csrfTokens[token]
	csrfMutex.RUnlock()

	if !exists {
		return false
	}

	if time.Now().After(expiry) {
		csrfMutex.Lock()
		delete(csrfTokens, token)
		csrfMutex.Unlock()
		return false
	}

	return true
}

// CSRFMiddleware validates CSRF tokens on POST requests
func CSRFMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			token := r.FormValue("csrf_token")
			if !ValidateCSRFToken(token) {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}
		}
		next(w, r)
	}
}
