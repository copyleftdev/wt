package fetcher

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-shiori/go-readability"
	"golang.org/x/net/html/charset"
)

// FetchAndExtractContent fetches the URL and extracts the main content using go-readability.
func FetchAndExtractContent(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Detect character encoding
	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	// Read the content
	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}

	// Parse the URL string into a *url.URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Parse content using go-readability
	article, err := readability.FromReader(strings.NewReader(string(data)), parsedURL)
	if err != nil {
		return "", err
	}

	return article.Content, nil
}
