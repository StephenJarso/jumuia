package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"jumuia/internal/db"
	"jumuia/internal/handlers"
	"jumuia/internal/middleware"
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
	"printf": fmt.Sprintf,
}

func main() {
	// Initialize database
	database := db.InitDB()
	defer database.Close()

	// Get port from environment variable (for Render deployment)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// Authentication routes (public)
	http.HandleFunc("/login", handlers.LoginHandler(database))
	http.HandleFunc("/register", handlers.RegisterHandler(database))
	http.HandleFunc("/logout", handlers.LogoutHandler())

	// Dashboard (public for now, can be protected later)
	http.HandleFunc("/dashboard", handlers.DashboardHandler(database))

	// Group routes (protected)
	http.HandleFunc("/groups/new", middleware.RequireAuth(handlers.NewGroupHandler(database)))
	http.HandleFunc("/groups/create", middleware.RequireAuth(handlers.CreateGroupHandler(database)))
	http.HandleFunc("/groups", handlers.ListGroupsHandler(database))
	http.HandleFunc("/groups/detail", middleware.RequireAuth(handlers.GroupDetailHandler(database)))
	http.HandleFunc("/groups/update-leader", middleware.RequireAuth(handlers.UpdateGroupLeaderHandler(database)))

	// Member routes (protected)
	http.HandleFunc("/members/new", middleware.RequireAuth(handlers.NewMemberHandler(database)))
	http.HandleFunc("/members/create", middleware.RequireAuth(handlers.CreateMemberHandler(database)))
	http.HandleFunc("/members", handlers.ListMembersHandler(database))
	http.HandleFunc("/members/account", middleware.RequireAuth(handlers.MemberAccountHandler(database)))

	// Season routes
	http.HandleFunc("/seasons/new", handlers.NewSeasonHandler(database))
	http.HandleFunc("/seasons/create", handlers.CreateSeasonHandler(database))
	http.HandleFunc("/seasons", handlers.ListSeasonsHandler(database))

	// Crop routes
	http.HandleFunc("/crops/new", handlers.NewCropHandler(database))
	http.HandleFunc("/crops/create", handlers.CreateCropHandler(database))
	http.HandleFunc("/crops", handlers.ListCropsHandler(database))

	// Loan routes (protected)
	http.HandleFunc("/loans/new", middleware.RequireAuth(handlers.NewLoanHandler(database)))
	http.HandleFunc("/loans/create", middleware.RequireAuth(handlers.CreateLoanHandler(database)))
	http.HandleFunc("/loans", handlers.ListLoansHandler(database))

	// Savings routes (protected)
	http.HandleFunc("/savings/new", middleware.RequireAuth(handlers.NewSavingsHandler(database)))
	http.HandleFunc("/savings/create", middleware.RequireAuth(handlers.CreateSavingsHandler(database)))
	http.HandleFunc("/savings", handlers.ListSavingsHandler(database))

	// Repayment routes (protected)
	http.HandleFunc("/repayments/new", middleware.RequireAuth(handlers.NewRepaymentHandler(database)))
	http.HandleFunc("/repayments/create", middleware.RequireAuth(handlers.CreateRepaymentHandler(database)))
	http.HandleFunc("/repayments", handlers.ListRepaymentsHandler(database))

	// Disaster routes
	http.HandleFunc("/disasters/new", handlers.NewDisasterHandler(database))
	http.HandleFunc("/disasters/create", handlers.CreateDisasterHandler(database))
	http.HandleFunc("/disasters", handlers.ListDisastersHandler(database))

	// Relief routes (protected)
	http.HandleFunc("/relief/new", middleware.RequireAuth(handlers.NewReliefHandler(database)))
	http.HandleFunc("/relief/create", middleware.RequireAuth(handlers.CreateReliefHandler(database)))
	http.HandleFunc("/relief", handlers.ListReliefHandler(database))

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
