package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// fetchURLContent fetches the content of a given URL and returns the root of the parsed HTML document.
func fetchURLContent(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching URL %s: %s", url, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return doc, nil
}

// extractText extracts and returns the human-readable text from the parsed HTML document.
func extractText(doc *goquery.Document) string {
	// Remove script and style elements
	doc.Find("script, style").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// Get the text content
	return doc.Text()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: wt <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	doc, err := fetchURLContent(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	plainText := extractText(doc)
	fmt.Println(plainText)
}
