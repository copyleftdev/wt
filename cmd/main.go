package main

import (
	"fmt"
	"os"

	"github.com/copyleftdev/wt/internal/config"
	"github.com/copyleftdev/wt/internal/fetcher"
	"github.com/copyleftdev/wt/internal/parser"
)

func main() {
	cfg, err := config.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	content, err := fetcher.FetchAndExtractContent(cfg.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching content: %v\n", err)
		os.Exit(1)
	}

	text, err := parser.ExtractTextFromHTML(content, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing content: %v\n", err)
		os.Exit(1)
	}

	// Output the extracted text
	fmt.Println(text)
}
