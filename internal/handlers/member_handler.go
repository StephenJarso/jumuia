package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"jumuia/internal/models"
	"jumuia/internal/repository"
)

// Show member form
func NewMemberHandler(w http.ResponseWriter, r *http.Request) {

	groupID := r.URL.Query().Get("group_id")

	tmpl := template.Must(template.ParseFiles("web/templates/member_form.html"))

	tmpl.Execute(w, groupID)
}

// Create member
func CreateMemberHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		groupIDStr := r.FormValue("group_id")

		groupID, _ := strconv.Atoi(groupIDStr)

		member := models.Member{
			GroupId: groupID,
			Name:    r.FormValue("name"),
			Phone:   r.FormValue("phone"),
			Role:    r.FormValue("role"),
		}

		_, err := repository.CreateMember(db, member)

		if err != nil {
			http.Error(w, "Error saving member", 500)
			return
		}

		http.Redirect(w, r, "/members?group_id="+groupIDStr, http.StatusSeeOther)
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

		tmpl := template.Must(template.ParseFiles("web/templates/member_list.html"))

		tmpl.Execute(w, members)
	}
}
