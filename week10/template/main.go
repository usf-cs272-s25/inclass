package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type SearchResult struct {
	Title string
	URL   string
}

func Search() []SearchResult {
	return []SearchResult{
		{"University of San Francisco", "https://www.usfca.edu"},
		{"New York Times", "http://www.nytimes.com"},
	}
}

func Serve() {
	tmpl, err := template.ParseFiles("demo.html")
	if err != nil {
		fmt.Println("template.ParseFiles: ", err)
	}

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results := Search()
		tmpl.Execute(w, results)
	})
	go http.ListenAndServe(":8080", nil)
}

func main() {
	index := MakeIndex()
	Serve(index)
	Crawl(index)
	for {
		time.Sleep(1 * time.Second)
	}
}
