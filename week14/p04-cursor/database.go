package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type SearchResult struct {
	URL     string
	Snippet string
	Score   float64
}

func initDatabase() (*sql.DB, error) {
	log.Println("Initializing database...")
	db, err := sql.Open("sqlite3", "search.db")
	if err != nil {
		return nil, err
	}

	// Create documents table
	log.Println("Creating documents table if not exists...")
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS documents (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT UNIQUE,
			text TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	// Create terms table for TF-IDF
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS terms (
			term TEXT,
			doc_id INTEGER,
			frequency INTEGER,
			PRIMARY KEY (term, doc_id),
			FOREIGN KEY (doc_id) REFERENCES documents(id)
		)
	`)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialization complete")
	return db, nil
}

func insertDocument(db *sql.DB, doc Document) error {
	log.Printf("Inserting document: %s", doc.URL)

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Insert document
	result, err := tx.Exec(`
		INSERT OR REPLACE INTO documents (url, text)
		VALUES (?, ?)
	`, doc.URL, doc.Text)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Get document ID
	docID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete existing terms for this document
	_, err = tx.Exec("DELETE FROM terms WHERE doc_id = ?", docID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Calculate term frequencies
	terms := strings.Fields(strings.ToLower(doc.Text))
	termFreq := make(map[string]int)
	for _, term := range terms {
		termFreq[term]++
	}

	// Insert term frequencies
	for term, freq := range termFreq {
		_, err = tx.Exec(`
			INSERT INTO terms (term, doc_id, frequency)
			VALUES (?, ?, ?)
		`, term, docID, freq)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	log.Printf("Successfully inserted document: %s", doc.URL)
	return tx.Commit()
}

func searchDocuments(db *sql.DB, query string) ([]SearchResult, error) {
	log.Printf("Searching for documents matching query: %q", query)

	// Split query into terms
	queryTerms := strings.Fields(strings.ToLower(query))
	if len(queryTerms) == 0 {
		return nil, nil
	}

	// Get total number of documents
	var totalDocs int
	err := db.QueryRow("SELECT COUNT(*) FROM documents").Scan(&totalDocs)
	if err != nil {
		return nil, err
	}

	// Build the query to calculate TF-IDF scores and extract snippets
	queryStr := `
		WITH doc_scores AS (
			SELECT 
				d.id,
				d.url,
				d.text,
				SUM(
					(t.frequency * 1.0 / (
						SELECT SUM(frequency) 
						FROM terms 
						WHERE doc_id = d.id
					)) * 
					(CAST(? AS FLOAT) / (
						SELECT COUNT(DISTINCT doc_id) 
						FROM terms 
						WHERE term = t.term
					))
				) as score
			FROM documents d
			JOIN terms t ON d.id = t.doc_id
			WHERE t.term IN (` + strings.Repeat("?,", len(queryTerms)-1) + "?)" + `
			GROUP BY d.id, d.url, d.text
			ORDER BY score DESC
			LIMIT 10
		)
		SELECT 
			url,
			substr(
				text,
				max(1, instr(lower(text), ?) - 50),
				100
			) as snippet,
			score
		FROM doc_scores
	`

	// Prepare arguments for the query
	args := make([]interface{}, len(queryTerms)+2)
	args[0] = totalDocs
	for i, term := range queryTerms {
		args[i+1] = term
	}
	args[len(args)-1] = queryTerms[0] // Use first term for snippet extraction

	// Execute query
	rows, err := db.Query(queryStr, args...)
	if err != nil {
		log.Printf("Error searching documents: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var result SearchResult
		if err := rows.Scan(&result.URL, &result.Snippet, &result.Score); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		// Clean up the snippet
		result.Snippet = strings.TrimSpace(result.Snippet)
		if len(result.Snippet) > 0 {
			results = append(results, result)
		}
	}

	log.Printf("Found %d documents matching query %q", len(results), query)
	return results, nil
}
