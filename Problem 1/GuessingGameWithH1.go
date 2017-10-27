// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 2 taken from https://data-representation.github.io/problems/go-web-applications.html
// Code adapted from https://golang.org/doc/articles/wiki/

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type","text/html");

	// Prints to screen
	fmt.Fprintln(w, "<h1>Guessing Game<h1>")
}

func main() {
	http.HandleFunc("/", requestHandler)

	// Sets the port to serve our web app on
	http.ListenAndServe(":8080", nil)
}