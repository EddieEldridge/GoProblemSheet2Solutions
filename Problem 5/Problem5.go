// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 5 taken from https://data-representation.github.io/problems/go-web-applications.html
// Code adapted from https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html


package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	
	tmpl := template.Must(template.ParseFiles("problem5.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}