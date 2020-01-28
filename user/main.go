package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Blen1234 dbname=users_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListUsers)
	http.HandleFunc("/signup.html", handleViewUser)
	http.HandleFunc("/login.html", handleViewUserLogin)
    http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/save", handleSaveUser)
	http.HandleFunc("/delete", handleDeleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
