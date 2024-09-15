package parser

import (
	"bytes"
	"strings"

	"github.com/copyleftdev/wt/internal/config"
	"golang.org/x/net/html"
)

var unwantedClassOrIDSubstrings = []string{
	"nav", "header", "footer", "sidebar", "menu", "account",
	"signin", "login", "signup", "register", "copyright",
	"subscribe", "breadcrumb", "advertisement", "ad-", "ads",
}

// ExtractTextFromHTML extracts text from the provided HTML content based on the configuration.
func ExtractTextFromHTML(htmlContent string, cfg *config.Config) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	extractText(doc, &buf, cfg)
	return buf.String(), nil
}

func extractText(n *html.Node, buf *bytes.Buffer, cfg *config.Config) {
	if n.Type == html.ElementNode {
		tagName := strings.ToLower(n.Data)
		if cfg.TagsConfig.ShouldIncludeTag(tagName) && !hasUnwantedAttributes(n) {
			text := getTextContent(n, cfg)
			text = strings.TrimSpace(text)
			if len(text) >= cfg.MinTextLength {
				buf.WriteString(text + "\n\n")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, buf, cfg)
	}
}

func getTextContent(n *html.Node, cfg *config.Config) string {
	var buf bytes.Buffer
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			tagName := strings.ToLower(c.Data)
			if cfg.TagsConfig.ShouldExcludeTag(tagName) || hasUnwantedAttributes(c) {
				continue
			}
		}
		buf.WriteString(getTextContent(c, cfg))
	}
	return buf.String()
}

func hasUnwantedAttributes(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, attr := range n.Attr {
		if attr.Key == "class" || attr.Key == "id" || attr.Key == "role" {
			value := strings.ToLower(attr.Val)
			for _, substr := range unwantedClassOrIDSubstrings {
				if strings.Contains(value, substr) {
					return true
				}
			}
		}
	}
	return false
}
