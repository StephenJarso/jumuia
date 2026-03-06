package handlers

import (
	"database/sql"
	"html/template"
	"jumuia/internal/models"
	"jumuia/internal/repository"
	"net/http"
)
// Show the form to create a new group
func NewGroupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/group.html")
	if err != nil {
		http.Error(w, "Error loading template", 500)
		return
	}
	tmpl.Execute(w, nil)
}
//Handle form submission and save group
func CreateGroupHandler(db *sql.DB)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		if r.Method != http.MethodPost{
			http.Redirect(w,r,"/groups/new",http.StatusSeeOther)
		}
		name:=r.FormValue("name")
		village := r.FormValue("village")
		district :=r.FormValue("district")
		 group:=models.Group{
			Name:name,
			Village: village,
			District: district,
		 }
		 _,err:=repository.CreateGroup(db,group)
		 if err != nil{
			http.Error(w,"Error saving group",500)
		 }
		 http.Redirect(w,r,"/groups",http.StatusSeeOther)
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

		tmpl, err := template.ParseFiles("web/templates/groups_list.html")
		if err != nil {
			http.Error(w, "Error loading template", 500)
			return
		}

		tmpl.Execute(w, groups)
	}
}