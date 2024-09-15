package tags

import (
	"reflect"
	"testing"
)

func TestNewTagsConfig_Defaults(t *testing.T) {
	tc := NewTagsConfig("", "")

	if len(tc.IncludeTags) == 0 {
		t.Error("Expected default include tags to be set")
	}

	if len(tc.ExcludeTags) == 0 {
		t.Error("Expected default exclude tags to be set")
	}
}

func TestNewTagsConfig_CustomTags(t *testing.T) {
	tc := NewTagsConfig("p,h1,h2", "div,span")

	expectedInclude := map[string]bool{"p": true, "h1": true, "h2": true}
	if !reflect.DeepEqual(tc.IncludeTags, expectedInclude) {
		t.Errorf("IncludeTags mismatch. Expected %v, got %v", expectedInclude, tc.IncludeTags)
	}

	expectedExclude := map[string]bool{"div": true, "span": true}
	if !reflect.DeepEqual(tc.ExcludeTags, expectedExclude) {
		t.Errorf("ExcludeTags mismatch. Expected %v, got %v", expectedExclude, tc.ExcludeTags)
	}
}

func TestTagsConfig_ShouldIncludeTag(t *testing.T) {
	tc := NewTagsConfig("p,h1,h2", "div,span")

	tests := []struct {
		tag      string
		expected bool
	}{
		{"p", true},
		{"h1", true},
		{"div", false},
		{"span", false},
		{"h3", false},
	}

	for _, test := range tests {
		result := tc.ShouldIncludeTag(test.tag)
		if result != test.expected {
			t.Errorf("ShouldIncludeTag(%s) = %v; want %v", test.tag, result, test.expected)
		}
	}
}

func TestTagsConfig_ShouldExcludeTag(t *testing.T) {
	tc := NewTagsConfig("", "script,style")

	tests := []struct {
		tag      string
		expected bool
	}{
		{"script", true},
		{"style", true},
		{"p", false},
		{"div", false},
	}

	for _, test := range tests {
		result := tc.ShouldExcludeTag(test.tag)
		if result != test.expected {
			t.Errorf("ShouldExcludeTag(%s) = %v; want %v", test.tag, result, test.expected)
		}
	}
}
