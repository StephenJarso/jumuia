package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

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
			Members []models.Member
		}{
			Members: members,
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
		name := r.FormValue("name")
		village := r.FormValue("village")
		district := r.FormValue("district")
		leaderIDStr := r.FormValue("leader_id")
		leaderID, _ := strconv.Atoi(leaderIDStr)

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

		data := struct {
			Group   *models.GroupWithLeader
			Members []models.Member
			Stats   map[string]interface{}
		}{
			Group:   group,
			Members: members,
			Stats:   stats,
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
