package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
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

// Show the form to create a new group
func NewGroupHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := repository.GetAllMembers(db)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}

		data := struct {
			Members   []models.Member
			CSRFToken string
		}{
			Members:   members,
			CSRFToken: middleware.GenerateCSRFToken(),
		}

		tmpl, err := template.New("group.html").Funcs(funcMap).ParseFiles("web/templates/group.html")
		if err != nil {
			http.Error(w, "Error loading templates", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Handle form submission and save group
func CreateGroupHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/groups/new", http.StatusSeeOther)
			return
		}

		// Validate required fields
		name := r.FormValue("name")
		village := r.FormValue("village")
		district := r.FormValue("district")
		leaderIDStr := r.FormValue("leader_id")

		name, ok := middleware.ValidateRequired(name, "name")
		if !ok {
			middleware.WriteValidationError(w, "Name is required")
			return
		}

		village, ok = middleware.ValidateRequired(village, "village")
		if !ok {
			middleware.WriteValidationError(w, "Village is required")
			return
		}

		district, ok = middleware.ValidateRequired(district, "district")
		if !ok {
			middleware.WriteValidationError(w, "District is required")
			return
		}

		leaderID, ok := middleware.ValidateInt(leaderIDStr, "leader_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid leader ID")
			return
		}

		group := models.Group{
			Name:     name,
			Village:  village,
			District: district,
			LeaderID: leaderID,
		}
		_, err := repository.CreateGroup(db, group)
		if err != nil {
			http.Error(w, "Error saving group", 500)
			return
		}
		http.Redirect(w, r, "/groups", http.StatusSeeOther)
	}
}

// List all groups
func ListGroupsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, err := repository.GetAllGroups(db)
		if err != nil {
			http.Error(w, "Error fetching groups", 500)
			return
		}

		tmpl, err := template.New("groups_list.html").Funcs(funcMap).ParseFiles("web/templates/groups_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, groups)
	}
}

// Show group detail page
func GroupDetailHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupIDStr := r.URL.Query().Get("id")
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			http.Error(w, "Invalid group ID", 400)
			return
		}

		group, err := repository.GetGroupByID(db, groupID)
		if err != nil {
			http.Error(w, "Group not found", 404)
			return
		}

		members, err := repository.GetMembersByGroup(db, groupID)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}

		stats, err := repository.GetGroupStats(db, groupID)
		if err != nil {
			http.Error(w, "Error loading statistics", 500)
			return
		}

		savings, err := repository.GetGroupSavings(db, groupID)
		if err != nil {
			http.Error(w, "Error loading savings", 500)
			return
		}

		relief, err := repository.GetGroupRelief(db, groupID)
		if err != nil {
			http.Error(w, "Error loading relief", 500)
			return
		}

		loans, err := repository.GetGroupLoans(db, groupID)
		if err != nil {
			http.Error(w, "Error loading loans", 500)
			return
		}

		savingsByMember, err := repository.GetGroupSavingsByMember(db, groupID)
		if err != nil {
			http.Error(w, "Error loading savings by member", 500)
			return
		}

		loansByStatus, err := repository.GetGroupLoansByStatus(db, groupID)
		if err != nil {
			http.Error(w, "Error loading loans by status", 500)
			return
		}

		// Convert to JSON for charts
		savingsByMemberJSON, _ := json.Marshal(savingsByMember)
		loansByStatusJSON, _ := json.Marshal(loansByStatus)

		data := struct {
			Group               *models.GroupWithLeader
			Members             []models.Member
			Stats               map[string]interface{}
			Savings             []map[string]interface{}
			Relief              []map[string]interface{}
			Loans               []map[string]interface{}
			SavingsByMember     []map[string]interface{}
			LoansByStatus       []map[string]interface{}
			SavingsByMemberJSON template.JS
			LoansByStatusJSON   template.JS
		}{
			Group:               group,
			Members:             members,
			Stats:               stats,
			Savings:             savings,
			Relief:              relief,
			Loans:               loans,
			SavingsByMember:     savingsByMember,
			LoansByStatus:       loansByStatus,
			SavingsByMemberJSON: template.JS(savingsByMemberJSON),
			LoansByStatusJSON:   template.JS(loansByStatusJSON),
		}

		tmpl, err := template.New("group_detail.html").Funcs(funcMap).ParseFiles("web/templates/group_detail.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, data)
	}
}

// Update group leader
func UpdateGroupLeaderHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/groups", http.StatusSeeOther)
			return
		}

		groupIDStr := r.FormValue("group_id")
		leaderIDStr := r.FormValue("leader_id")

		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			http.Error(w, "Invalid group ID", 400)
			return
		}

		leaderID, err := strconv.Atoi(leaderIDStr)
		if err != nil {
			http.Error(w, "Invalid leader ID", 400)
			return
		}

		err = repository.UpdateGroupLeader(db, groupID, leaderID)
		if err != nil {
			http.Error(w, "Error updating leader", 500)
			return
		}

		http.Redirect(w, r, "/groups/detail?id="+groupIDStr, http.StatusSeeOther)
	}
}
