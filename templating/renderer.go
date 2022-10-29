package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}

type PostRenderer struct{
  templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
  templ, err := template.ParseFS(postTemplates, "templates/*.gohtml") 

  if err != nil {
    return nil, err
  }

  return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
  
  if err := r.templ.Execute(w,p); err != nil {
    return err
  }

  return nil
}
