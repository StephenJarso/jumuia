package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"jumuia/internal/db"
	"jumuia/internal/handlers"
)

// Template functions
var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"slice": func(s string, i int) string {
		if i >= len(s) {
			return s
		}
		return s[:i]
	},
}

func main() {
	// Initialize database
	database := db.InitDB()
	defer database.Close()

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// Dashboard
	http.HandleFunc("/dashboard", handlers.DashboardHandler(database))

	// Group routes
	http.HandleFunc("/groups/new", handlers.NewGroupHandler(database))
	http.HandleFunc("/groups/create", handlers.CreateGroupHandler(database))
	http.HandleFunc("/groups", handlers.ListGroupsHandler(database))
	http.HandleFunc("/groups/detail", handlers.GroupDetailHandler(database))
	http.HandleFunc("/groups/update-leader", handlers.UpdateGroupLeaderHandler(database))

	// Member routes
	http.HandleFunc("/members/new", handlers.NewMemberHandler(database))
	http.HandleFunc("/members/create", handlers.CreateMemberHandler(database))
	http.HandleFunc("/members", handlers.ListMembersHandler(database))
	http.HandleFunc("/members/account", handlers.MemberAccountHandler(database))

	// Season routes
	http.HandleFunc("/seasons/new", handlers.NewSeasonHandler)
	http.HandleFunc("/seasons/create", handlers.CreateSeasonHandler(database))
	http.HandleFunc("/seasons", handlers.ListSeasonsHandler(database))

	// Crop routes
	http.HandleFunc("/crops/new", handlers.NewCropHandler)
	http.HandleFunc("/crops/create", handlers.CreateCropHandler(database))
	http.HandleFunc("/crops", handlers.ListCropsHandler(database))

	// Loan routes
	http.HandleFunc("/loans/new", handlers.NewLoanHandler(database))
	http.HandleFunc("/loans/create", handlers.CreateLoanHandler(database))
	http.HandleFunc("/loans", handlers.ListLoansHandler(database))

	// Savings routes
	http.HandleFunc("/savings/new", handlers.NewSavingsHandler(database))
	http.HandleFunc("/savings/create", handlers.CreateSavingsHandler(database))
	http.HandleFunc("/savings", handlers.ListSavingsHandler(database))

	// Repayment routes
	http.HandleFunc("/repayments/new", handlers.NewRepaymentHandler(database))
	http.HandleFunc("/repayments/create", handlers.CreateRepaymentHandler(database))
	http.HandleFunc("/repayments", handlers.ListRepaymentsHandler(database))

	// Disaster routes
	http.HandleFunc("/disasters/new", handlers.NewDisasterHandler)
	http.HandleFunc("/disasters/create", handlers.CreateDisasterHandler(database))
	http.HandleFunc("/disasters", handlers.ListDisastersHandler(database))

	// Relief routes
	http.HandleFunc("/relief/new", handlers.NewReliefHandler(database))
	http.HandleFunc("/relief/create", handlers.CreateReliefHandler(database))
	http.HandleFunc("/relief", handlers.ListReliefHandler(database))

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
