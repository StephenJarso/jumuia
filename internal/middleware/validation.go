package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

// ValidateRequired checks if a form value is non-empty
func ValidateRequired(value, fieldName string) (string, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", false
	}
	return value, true
}

// ValidateInt parses and validates an integer value
func ValidateInt(value, fieldName string) (int, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return num, true
}

// ValidateFloat parses and validates a float value
func ValidateFloat(value, fieldName string) (float64, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

// ValidatePhone validates phone number format (basic)
func ValidatePhone(value string) (string, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", true // Phone is optional
	}
	// Basic validation: starts with + and has 10-15 digits
	if !strings.HasPrefix(value, "+") {
		return "", false
	}
	digits := strings.TrimPrefix(value, "+")
	if len(digits) < 10 || len(digits) > 15 {
		return "", false
	}
	return value, true
}

// WriteValidationError writes a validation error response
func WriteValidationError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}
