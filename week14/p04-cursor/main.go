package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Parse command line arguments
	seedURL := flag.String("url", "", "Seed URL to start crawling from")
	port := flag.String("port", "8080", "Port to run the web server on")
	flag.Parse()

	if *seedURL == "" {
		fmt.Println("Please provide a seed URL using -url flag")
		os.Exit(1)
	}

	// Initialize database
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create channels for communication
	urlChan := make(chan string, 100)
	docChan := make(chan Document, 100)
	doneChan := make(chan bool)

	// Start the crawler
	go startCrawler(*seedURL, urlChan, docChan, doneChan)

	// Start the web server
	go startWebServer(*port, db)

	// Process documents and insert them into the database
	go func() {
		for doc := range docChan {
			if err := insertDocument(db, doc); err != nil {
				log.Printf("Error inserting document: %v", err)
			}
		}
	}()

	// Wait for crawler to finish
	<-doneChan
	log.Println("Crawling completed")
}
