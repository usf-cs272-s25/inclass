package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

type PageData struct {
	Numbers []int
	Error   string
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, PageData{})
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		input := r.FormValue("numbers")
		data := PageData{}

		// Split the input string into individual numbers
		numberStrings := strings.Fields(input)
		if len(numberStrings) == 0 {
			data.Error = "No numbers provided"
			tmpl := template.Must(template.ParseFiles("static/index.html"))
			tmpl.Execute(w, data)
			return
		}

		// Convert strings to integers
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				data.Error = fmt.Sprintf("Error: '%s' is not a valid integer", numStr)
				tmpl := template.Must(template.ParseFiles("static/index.html"))
				tmpl.Execute(w, data)
				return
			}
			numbers[i] = num
		}

		// Sort the numbers
		sort.Ints(numbers)
		data.Numbers = numbers

		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, data)
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the main route
	http.HandleFunc("/", sortHandler)

	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Try different ports
	ports := []int{8080, 8081, 8082, 8083}
	var server *http.Server
	var listener net.Listener
	var err error

	for _, port := range ports {
		addr := fmt.Sprintf(":%d", port)
		listener, err = net.Listen("tcp", addr)
		if err == nil {
			server = &http.Server{Addr: addr}
			fmt.Printf("Server starting on http://localhost:%d\n", port)
			break
		}
	}

	if err != nil {
		fmt.Printf("Failed to start server on any port: %v\n", err)
		os.Exit(1)
	}

	// Start server in a goroutine
	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nShutting down server...")
}
