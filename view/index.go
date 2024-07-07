package view

import (
	"html/template"
	"log"
)

type pageData struct {
	Title   string
	Content string
}

func (s *Static) pageIndex(tmpl *template.Template) {

	f := s.createFile("/index.html")

	page := pageData{
		Title:   "Man on Box",
		Content: "Working on it...",
	}

	if err := tmpl.ExecuteTemplate(f, "homepage", page); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

}
