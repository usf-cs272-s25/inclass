package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type SearchResponse struct {
	Title string
	URL   string
}

func Search() []SearchResponse {
	return []SearchResponse{
		{"Google", "https://www.google.com"},
		{"University of San Francisco", "https://www.usfca.edu"},
	}
}

func Serve() {
	tmpl, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Println("Failed to parse template")
	}

	http.Handle("/search", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		hits := Search()
		tmpl.Execute(w, hits)
	}))

	go http.ListenAndServe(":8080", nil)
}

func main() {
	Serve()
	for {
		time.Sleep(1 * time.Second)
	}
}
