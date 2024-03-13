package main

import (
	"github.com/google/uuid"
	"reflect"
	"slices"
	"time"
)

type User struct {
	Id        string
	Name      string
	CreatedAt time.Time
}

var issueLabels []string = []string{"bug", "improvement"}
var issueStatus []string = []string{"open", "closed"}

type Issue struct {
	Id        string
	Title     string
	Content   string
	Creator   User
	Labels    string
	Status    string
	CreatedAt time.Time
}

var users []User
var issues []Issue

func init() {
	defaultUsers := []map[string]string{
		{"name": "Ozgur A"},
	}
	defaultIssues := []map[string]string{
		{"title": "button is not clickable", "content": "fix this", "labels": "bug"},
		{"title": "improve button readability", "content": "please add this", "labels": "improvement"},
	}
	for _, d := range defaultUsers {
		createUser(d)
	}
	for _, d := range defaultIssues {
		createIssue(d)
	}
}

func getIssue(issueId string) Issue {
	idx := slices.IndexFunc(issues, func(issue Issue) bool {
		return issue.Id == issueId
	})
	return issues[idx]
}

func createIssue(issueData map[string]string) Issue {
	issue := Issue{
		Id:        uuid.NewString(),
		Title:     issueData["title"],
		Content:   issueData["content"],
		Creator:   users[0],
		Labels:    issueData["labels"],
		Status:    "open",
		CreatedAt: time.Now(),
	}
	issues = append(issues, issue)
	return issue
}

func editIssue(issueData map[string]string) {
	idx := slices.IndexFunc(issues, func(issue Issue) bool {
		return issue.Id == issueData["Id"]
	})
	delete(issueData, "Id")
	issue := &issues[idx]
	for key, value := range issueData {
		field := reflect.ValueOf(issue).Elem().FieldByName(key)
		field.SetString(value)
	}
}

func createUser(userData map[string]string) User {
	user := User{
		Id:        uuid.NewString(),
		Name:      userData["name"],
		CreatedAt: time.Now(),
	}
	users = append(users, user)
	return user
}
