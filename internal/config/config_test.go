package config

import (
	"reflect"
	"testing"
)

func TestParseFlags_Defaults(t *testing.T) {
	args := []string{"https://example.com"}
	cfg, err := ParseFlags(args)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.URL != "https://example.com" {
		t.Errorf("Expected URL to be 'https://example.com', got '%s'", cfg.URL)
	}

	if len(cfg.TagsConfig.IncludeTags) == 0 {
		t.Error("Expected default include tags to be set")
	}

	if len(cfg.TagsConfig.ExcludeTags) == 0 {
		t.Error("Expected default exclude tags to be set")
	}
}

func TestParseFlags_CustomTags(t *testing.T) {
	args := []string{"-include", "p,h1,h2", "-exclude", "div,span", "https://example.com"}
	cfg, err := ParseFlags(args)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedInclude := map[string]bool{"p": true, "h1": true, "h2": true}
	if !reflect.DeepEqual(cfg.TagsConfig.IncludeTags, expectedInclude) {
		t.Errorf("IncludeTags mismatch. Expected %v, got %v", expectedInclude, cfg.TagsConfig.IncludeTags)
	}

	expectedExclude := map[string]bool{"div": true, "span": true}
	if !reflect.DeepEqual(cfg.TagsConfig.ExcludeTags, expectedExclude) {
		t.Errorf("ExcludeTags mismatch. Expected %v, got %v", expectedExclude, cfg.TagsConfig.ExcludeTags)
	}
}

func TestParseFlags_MissingURL(t *testing.T) {
	args := []string{}
	_, err := ParseFlags(args)
	if err == nil {
		t.Fatal("Expected an error when URL is missing")
	}
}
