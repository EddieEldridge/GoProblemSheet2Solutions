// Go Code written by Edward Eldridge (G00337490)
// GMIT - Software Development
// Problem 4 taken from https://data-representation.github.io/problems/go-web-applications.html
// Code adapted from https://getbootstrap.com/docs/4.0/getting-started/introduction/#starter-template and https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server


package main

import (
	//"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Browser renders html tags
	w.Header().Set("Content-Type","text/html")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./guess")))
	http.ListenAndServe(":8080", nil)
}