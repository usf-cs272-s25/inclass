package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type SortResponse struct {
	Numbers []int `json:"numbers"`
}

func handleSort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get numbers from form data
	numbersStr := r.FormValue("numbers")
	if numbersStr == "" {
		http.Error(w, "No numbers provided", http.StatusBadRequest)
		return
	}

	// Split and convert to integers
	numbers := make([]int, 0)
	for _, numStr := range strings.Fields(numbersStr) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid number: %s", numStr), http.StatusBadRequest)
			return
		}
		numbers = append(numbers, num)
	}

	// Sort the numbers
	sort.Ints(numbers)

	// Check if this is an API request
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SortResponse{Numbers: numbers})
		return
	}

	// For web requests, return to the index page with the sorted numbers
	http.Redirect(w, r, fmt.Sprintf("/?numbers=%s&sorted=%v", numbersStr, numbers), http.StatusSeeOther)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get query parameters
	numbers := r.URL.Query().Get("numbers")
	sorted := r.URL.Query().Get("sorted")

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Number Sorter</title>
    <link rel="stylesheet" href="/css/styles.css">
</head>
<body>
    <div class="container">
        <h1>Number Sorter</h1>
        <form action="/sort" method="post">
            <div class="form-group">
                <label for="numbers">Enter numbers (space-separated):</label>
                <input type="text" id="numbers" name="numbers" value="%s" placeholder="5 2 8 1 3" required>
            </div>
            <button type="submit">Sort Numbers</button>
        </form>

        %s
    </div>
</body>
</html>`, numbers, func() string {
		if sorted != "" {
			return fmt.Sprintf(`<div class="result">
                <h2>Sorted Numbers:</h2>
                <div class="numbers">%s</div>
            </div>`, sorted)
		}
		return ""
	}())
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Handle routes
	http.HandleFunc("/sort", handleSort)
	http.HandleFunc("/index", handleIndex)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
