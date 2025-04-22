package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Document struct {
	URL  string
	Text string
}

var (
	visitedURLs = make(map[string]bool)
	mu          sync.Mutex
	// Tags to skip when extracting text
	skipTags = map[string]bool{
		"style":    true,
		"script":   true,
		"head":     true,
		"meta":     true,
		"link":     true,
		"noscript": true,
		"iframe":   true,
		"svg":      true,
	}
)

func startCrawler(seedURL string, urlChan chan string, docChan chan Document, doneChan chan bool) {
	log.Printf("Starting crawler with seed URL: %s", seedURL)

	// Create a wait group to track all crawling goroutines
	var wg sync.WaitGroup

	// Start with the seed URL
	urlChan <- seedURL

	// Start multiple worker goroutines to process URLs
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for urlStr := range urlChan {
				mu.Lock()
				if visitedURLs[urlStr] {
					mu.Unlock()
					continue
				}
				visitedURLs[urlStr] = true
				mu.Unlock()

				log.Printf("Crawling URL: %s", urlStr)

				// Add a small delay between requests to be polite
				time.Sleep(100 * time.Millisecond)

				// Fetch and parse the page
				resp, err := http.Get(urlStr)
				if err != nil {
					log.Printf("Error fetching %s: %v", urlStr, err)
					continue
				}
				defer resp.Body.Close()

				// Only process HTML content
				if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
					log.Printf("Skipping non-HTML content: %s", urlStr)
					continue
				}

				// Parse HTML
				doc, err := html.Parse(resp.Body)
				if err != nil {
					log.Printf("Error parsing HTML from %s: %v", urlStr, err)
					continue
				}

				// Extract text and links
				text := extractText(doc)
				log.Printf("Extracted %d characters of text from %s", len(text), urlStr)
				docChan <- Document{URL: urlStr, Text: text}

				// Extract and queue new URLs
				linkCount := extractLinks(doc, urlStr, urlChan)
				log.Printf("Found %d links on %s", linkCount, urlStr)
			}
		}()
	}

	// Wait for all workers to finish
	wg.Wait()
	log.Println("Crawler finished processing all URLs")
	doneChan <- true
}

func extractText(n *html.Node) string {
	var text strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		// Skip if this is a tag we want to ignore
		if n.Type == html.ElementNode && skipTags[n.Data] {
			return
		}

		// Only process text nodes
		if n.Type == html.TextNode {
			// Clean up the text
			cleanText := strings.TrimSpace(n.Data)
			if cleanText != "" {
				text.WriteString(cleanText)
				text.WriteString(" ")
			}
		}

		// Process children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Find the body element
	var body *html.Node
	var findBody func(*html.Node)
	findBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findBody(c)
		}
	}
	findBody(n)

	// If we found the body, extract text from it
	if body != nil {
		f(body)
	} else {
		// Fallback to processing the whole document
		f(n)
	}

	return strings.TrimSpace(text.String())
}

func extractLinks(n *html.Node, baseURL string, urlChan chan string) int {
	linkCount := 0
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					base, err := url.Parse(baseURL)
					if err != nil {
						continue
					}
					absoluteURL := base.ResolveReference(link).String()

					// Only follow links from the same domain
					if strings.Contains(absoluteURL, base.Host) {
						urlChan <- absoluteURL
						linkCount++
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return linkCount
}
