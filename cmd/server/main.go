package main

import (
	"fmt"
	"jumuia/internal/db"
	"net/http"
	"log"
)

func main(){
	db := db.InitDB()
	defer db.Close()
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Println(w,"Welcome to jumuia")
	})
	log.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}