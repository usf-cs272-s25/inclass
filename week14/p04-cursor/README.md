# Go Search Engine

A simple search engine built in Go that crawls websites and provides a web interface for searching through the crawled content.

## Features

- Web crawler that starts from a seed URL
- Extracts user-visible text from HTML documents
- Stores content in SQLite database with full-text search
- Web interface for searching through crawled content
- Concurrent crawling and web serving using Go channels

## Requirements

- Go 1.21 or later
- SQLite3

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

Run the search engine with a seed URL:

```bash
go run . -url https://example.com -port 8080
```

- `-url`: The seed URL to start crawling from (required)
- `-port`: Port to run the web server on (default: 8080)

## How it Works

1. The crawler starts from the seed URL and recursively follows links
2. For each page, it extracts the visible text content
3. The content is stored in a SQLite database with full-text search capabilities
4. The web server provides an interface to search through the crawled content
5. Search results are ranked by relevance using SQLite's FTS5

## Project Structure

- `main.go`: Program entry point and coordination
- `crawler.go`: Web crawler implementation
- `database.go`: SQLite database operations
- `webserver.go`: Web server implementation
- `templates/index.html`: Search interface template 