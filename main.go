package main

import (
	"log"
	"net/http"
)


func main() {
	PORT := "8000"

	handler := Handler{}

	// Client Routes
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("GET /", handler.Client.Issue.GetIssues)
	http.HandleFunc("GET /issue/new", handler.Client.Issue.PostNewIssue)
	http.HandleFunc("GET /issue/{id}", handler.Client.Issue.GetIssueByID)

	// Server Routes
	http.HandleFunc("POST /api/issue", handler.Server.Issue.PostNewIssue)
	http.HandleFunc("PUT /api/issue", handler.Server.Issue.PutNewIssue)

	// Start server
	log.Println("Listening on http://localhost:" + PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
