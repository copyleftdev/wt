package config

import (
	"flag"
	"fmt"

	"github.com/copyleftdev/wt/internal/tags"
)

type Config struct {
	URL           string
	TagsConfig    *tags.TagsConfig
	MinTextLength int
}

func ParseFlags(args []string) (*Config, error) {
	fs := flag.NewFlagSet("wt", flag.ContinueOnError)

	includeTags := fs.String("include", "", "Comma-separated list of tags to include (default: common content tags)")
	excludeTags := fs.String("exclude", "", "Comma-separated list of tags to exclude (default: non-content tags)")
	minTextLength := fs.Int("minlength", 50, "Minimum length of text blocks to include")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if fs.NArg() < 1 {
		fmt.Println("Usage: wt [options] <URL>")
		fs.PrintDefaults()
		return nil, fmt.Errorf("URL is required")
	}

	tagsConfig := tags.NewTagsConfig(*includeTags, *excludeTags)

	cfg := &Config{
		URL:           fs.Arg(0),
		TagsConfig:    tagsConfig,
		MinTextLength: *minTextLength,
	}

	return cfg, nil
}
