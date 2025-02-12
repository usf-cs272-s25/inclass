package main

import (
	"fmt"
	"log"
	"net/http"
)

func registerHandlers() {
	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Search results go here")
	})

}
func main() {
	registerHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
