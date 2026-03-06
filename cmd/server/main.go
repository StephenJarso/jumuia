package main

import (
	_"fmt"
	"jumuia/internal/db"
	"jumuia/internal/handlers"
	"log"
	"net/http"
)

func main(){
	db := db.InitDB()
	defer db.Close()
	// http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
	// 	fmt.Println(w,"Welcome to jumuia")
	// })
	http.HandleFunc("/groups/new",handlers.NewGroupHandler)
	http.HandleFunc("/groups/create",handlers.CreateGroupHandler(db))
	http.HandleFunc("/groups",handlers.ListGroupsHandler(db))
	log.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}