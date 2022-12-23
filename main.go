package main

import (
	"cmtk/auth"
	"net/http"
)

func main() {
	http.HandleFunc("/api/login", auth.Login)
	//http.HandleFunc("/api/recentRecords", query.Recentrecords)
	http.ListenAndServe(":8080", nil)
}
