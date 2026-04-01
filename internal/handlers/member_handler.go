package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"jumuia/internal/middleware"
	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Template functions
var memberFuncMap = template.FuncMap{
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

// Show member form
func NewMemberHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, err := repository.GetAllGroups(db)
		if err != nil {
			http.Error(w, "Error loading groups", 500)
			return
		}

		data := struct {
			Groups    []models.GroupWithLeader
			CSRFToken string
		}{
			Groups:    groups,
			CSRFToken: middleware.GenerateCSRFToken(),
		}

		tmpl, err := template.New("member_form.html").Funcs(memberFuncMap).ParseFiles("web/templates/member_form.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}
		tmpl.Execute(w, data)
	}
}

// Create member
func CreateMemberHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Validate required fields
		groupIDStr := r.FormValue("group_id")
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		role := r.FormValue("role")

		groupID, ok := middleware.ValidateInt(groupIDStr, "group_id")
		if !ok {
			middleware.WriteValidationError(w, "Invalid group ID")
			return
		}

		name, ok = middleware.ValidateRequired(name, "name")
		if !ok {
			middleware.WriteValidationError(w, "Name is required")
			return
		}

		phone, ok = middleware.ValidatePhone(phone)
		if !ok {
			middleware.WriteValidationError(w, "Invalid phone number")
			return
		}

		role, ok = middleware.ValidateRequired(role, "role")
		if !ok {
			middleware.WriteValidationError(w, "Role is required")
			return
		}

		member := models.Member{
			GroupId: groupID,
			Name:    name,
			Phone:   phone,
			Role:    role,
		}

		_, err := repository.CreateMember(db, member)
		if err != nil {
			http.Error(w, "Error saving member", 500)
			return
		}

		http.Redirect(w, r, "/groups/detail?id="+groupIDStr, http.StatusSeeOther)
	}
}

// List members
func ListMembersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupIDStr := r.URL.Query().Get("group_id")
		groupID, _ := strconv.Atoi(groupIDStr)

		members, err := repository.GetMembersByGroup(db, groupID)
		if err != nil {
			http.Error(w, "Error loading members", 500)
			return
		}

		tmpl, err := template.New("member_list.html").Funcs(memberFuncMap).ParseFiles("web/templates/member_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}
		tmpl.Execute(w, members)
	}
}

// Show member account page
func MemberAccountHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberIDStr := r.URL.Query().Get("id")
		memberID, err := strconv.Atoi(memberIDStr)
		if err != nil {
			http.Error(w, "Invalid member ID", 400)
			return
		}

		member, err := repository.GetMemberByID(db, memberID)
		if err != nil {
			http.Error(w, "Member not found", 404)
			return
		}

		group, err := repository.GetGroupByID(db, member.GroupId)
		if err != nil {
			http.Error(w, "Error loading group", 500)
			return
		}

		summary, err := repository.GetMemberFinancialSummary(db, memberID)
		if err != nil {
			http.Error(w, "Error loading financial summary", 500)
			return
		}

		data := struct {
			Member  *models.Member
			Group   *models.GroupWithLeader
			Summary map[string]interface{}
		}{
			Member:  member,
			Group:   group,
			Summary: summary,
		}

		tmpl, err := template.New("member_account.html").Funcs(memberFuncMap).ParseFiles("web/templates/member_account.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, data)
	}
}
