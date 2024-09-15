package tags

import (
	"strings"
)

var (
	defaultIncludeTags = []string{
		"article", "section", "p", "h1", "h2", "h3", "h4", "h5", "h6",
		"li", "blockquote", "pre", "code", "td", "th",
	}
	defaultExcludeTags = []string{
		"script", "style", "nav", "footer", "header", "aside", "form",
		"noscript", "iframe", "input", "button", "select", "label",
		"option", "link", "meta", "figure", "figcaption", "dialog",
		"menu", "svg", "canvas", "video", "audio", "embed", "object",
		"applet", "area", "base", "basefont", "bgsound", "blink",
		"col", "colgroup", "command", "content", "data", "datalist",
		"del", "details", "dir", "fieldset", "font", "frame", "frameset",
		"head", "img", "input", "isindex", "keygen", "marquee",
		"menuitem", "meter", "noembed", "noframes", "optgroup", "output",
		"param", "picture", "progress", "rp", "rt", "ruby", "s", "samp",
		"section", "select", "small", "source", "strike", "style",
		"sub", "summary", "sup", "template", "textarea", "track",
		"var", "wbr", "xmp",
	}
)

type TagsConfig struct {
	IncludeTags map[string]bool
	ExcludeTags map[string]bool
}

func NewTagsConfig(includeTagsStr, excludeTagsStr string) *TagsConfig {
	includeTags := parseTagList(includeTagsStr)
	excludeTags := parseTagList(excludeTagsStr)

	if len(includeTags) == 0 {
		includeTags = make(map[string]bool)
		for _, tag := range defaultIncludeTags {
			includeTags[tag] = true
		}
	}

	if len(excludeTags) == 0 {
		excludeTags = make(map[string]bool)
		for _, tag := range defaultExcludeTags {
			excludeTags[tag] = true
		}
	}

	return &TagsConfig{
		IncludeTags: includeTags,
		ExcludeTags: excludeTags,
	}
}

func parseTagList(tagsStr string) map[string]bool {
	tagsMap := make(map[string]bool)
	if tagsStr == "" {
		return tagsMap
	}
	tags := strings.Split(tagsStr, ",")
	for _, tag := range tags {
		tag = strings.TrimSpace(strings.ToLower(tag))
		if tag != "" {
			tagsMap[tag] = true
		}
	}
	return tagsMap
}

func (tc *TagsConfig) ShouldIncludeTag(tag string) bool {
	tag = strings.ToLower(tag)
	if tc.ExcludeTags[tag] {
		return false
	}
	if tc.IncludeTags[tag] {
		return true
	}
	return false
}

func (tc *TagsConfig) ShouldExcludeTag(tag string) bool {
	tag = strings.ToLower(tag)
	return tc.ExcludeTags[tag]
}
