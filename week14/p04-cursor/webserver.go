package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type SearchResponse struct {
	Query   string
	Results []SearchResult
}

func startWebServer(port string, db *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleSearch(w, r, db)
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		handleSearch(w, r, db)
	})

	log.Printf("Starting web server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting web server: %v", err)
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "GET" && r.URL.Path == "/" {
		renderTemplate(w, "index.html", nil)
		return
	}

	query := r.FormValue("q")
	if query == "" {
		renderTemplate(w, "index.html", nil)
		return
	}

	results, err := searchDocuments(db, query)
	if err != nil {
		http.Error(w, "Error searching documents", http.StatusInternalServerError)
		return
	}

	// Highlight matching terms in snippets
	for i := range results {
		for _, term := range strings.Fields(strings.ToLower(query)) {
			results[i].Snippet = strings.ReplaceAll(
				results[i].Snippet,
				term,
				"<strong>"+term+"</strong>",
			)
		}
	}

	response := SearchResponse{
		Query:   query,
		Results: results,
	}

	renderTemplate(w, "index.html", response)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
