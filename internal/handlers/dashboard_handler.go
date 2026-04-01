package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"jumuia/internal/repository"
)

// DashboardHandler shows the main dashboard with summary statistics
func DashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get counts for dashboard
		groups, err := repository.GetAllGroups(db)
		if err != nil {
			http.Error(w, "Error fetching groups", 500)
			return
		}

		members, err := repository.GetAllMembers(db)
		if err != nil {
			http.Error(w, "Error fetching members", 500)
			return
		}

		loans, err := repository.GetAllLoans(db)
		if err != nil {
			http.Error(w, "Error fetching loans", 500)
			return
		}

		savings, err := repository.GetAllSavings(db)
		if err != nil {
			http.Error(w, "Error fetching savings", 500)
			return
		}

		disasters, err := repository.GetAllDisasters(db)
		if err != nil {
			http.Error(w, "Error fetching disasters", 500)
			return
		}

		// Calculate totals
		totalSavings := 0.0
		for _, s := range savings {
			totalSavings += s.Amount
		}

		totalLoans := 0.0
		activeLoans := 0
		for _, l := range loans {
			totalLoans += l.Amount
			if l.Status == "active" {
				activeLoans++
			}
		}

		// Calculate total repayments
		totalRepayments := 0.0
		repayments, err := repository.GetAllRepayments(db)
		if err == nil {
			for _, rep := range repayments {
				totalRepayments += rep.Amount
			}
		}

		data := struct {
			GroupCount      int
			MemberCount     int
			LoanCount       int
			ActiveLoans     int
			SavingsCount    int
			DisasterCount   int
			TotalSavings    float64
			TotalLoans      float64
			TotalRepayments float64
			Groups          []interface{}
		}{
			GroupCount:      len(groups),
			MemberCount:     len(members),
			LoanCount:       len(loans),
			ActiveLoans:     activeLoans,
			SavingsCount:    len(savings),
			DisasterCount:   len(disasters),
			TotalSavings:    totalSavings,
			TotalLoans:      totalLoans,
			TotalRepayments: totalRepayments,
		}

		tmpl, err := template.ParseFiles("web/templates/dashboard.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, data)
	}
}
