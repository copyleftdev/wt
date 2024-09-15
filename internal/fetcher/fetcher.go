package fetcher

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-shiori/go-readability"
	"golang.org/x/net/html/charset"
)

func FetchAndExtractContent(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyReader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	article, err := readability.FromReader(strings.NewReader(string(data)), parsedURL)
	if err != nil {
		return "", err
	}

	return article.Content, nil
}
