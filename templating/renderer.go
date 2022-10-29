package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// PostRenderer renders data into HTML
type PostRenderer struct {
	templ *template.Template
}

// NewPostRenderer creates a new PostRenderer
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

// Render renders post into HTML
func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

// RenderIndex creates an HTML index page given a collection of posts
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
  return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}
