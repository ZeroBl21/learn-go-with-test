package blogrenderer

import "strings"

// Post is a representation of a post
type Post struct {
	Title, Body, Description string
	Tags                     []string
}

// SanitisedTitle returns the title of the post with spaces replaced by dashes for pleasant URLs
func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

