package main

import (
	"fmt"
	"log"
	"net/http"
)

func registerHandlers(idx InvIndex) {
	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		// go lookup search term in index data structure
		// write results into w (ResponseWriter)

		v, ok := idx[searchTerm]
		fmt.Fprintf(w, "Search results go here")
	})

}
func main() {
	registerHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
