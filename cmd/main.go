package main

import (
	"fmt"
	"os"

	"github.com/copyleftdev/wt/internal/config"
	"github.com/copyleftdev/wt/internal/fetcher"
	"github.com/copyleftdev/wt/internal/parser"
)

func main() {
	// Parse command-line arguments and flags
	cfg, err := config.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Fetch and extract the main content using go-readability
	content, err := fetcher.FetchAndExtractContent(cfg.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching content: %v\n", err)
		os.Exit(1)
	}

	// Parse the extracted content and apply tag and attribute filtering
	text, err := parser.ExtractTextFromHTML(content, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing content: %v\n", err)
		os.Exit(1)
	}

	// Output the extracted text
	fmt.Println(text)
}
