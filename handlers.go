package main

import (
	"github.com/a-h/templ"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	Server ServerHandler
	Client ClientHandler
}

type ServerHandler struct {
	Issue IssueHandlerServer
}

type ClientHandler struct {
	Issue IssueHandlerClient
}

type IssueHandlerServer struct{}
type IssueHandlerClient struct{}

func (ih IssueHandlerClient) GetIssues(w http.ResponseWriter, r *http.Request) {
	templ.Handler(index(issues)).ServeHTTP(w, r)
}
func (ih IssueHandlerClient) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	issue := getIssue(id)
	templ.Handler(IssueDetail(issue)).ServeHTTP(w, r)
}
func (ih IssueHandlerClient) PostNewIssue(w http.ResponseWriter, r *http.Request) {
	templ.Handler(AddIssue()).ServeHTTP(w, r)
}

func (ih IssueHandlerServer) PostNewIssue(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	createIssue(map[string]string{"title": r.Form["title"][0], "content": r.Form["content"][0], "labels": strings.Join(r.Form["labels"], ",")})
	w.Header().Add("HX-Redirect", "/")
}
func (ih IssueHandlerServer) PutNewIssue(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := map[string]string{}
	for key, valueArr := range r.Form {
		data[key] = valueArr[0]
	}
	editIssue(data)
	w.Header().Add("HX-Redirect", "/")
}
