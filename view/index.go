package view

import (
	"html/template"
)

type pageData struct {
	Title   string
	Content string
}

func (s *Static) pageIndex() {

	f := s.createFile("/index.html")

	tmpl := template.Must(template.ParseFiles(
		"./view/templates/base.html",
		"./view/templates/index.html",
	))
	tmpl.Execute(f, pageData{
		Title:   "Man on Box",
		Content: "Working on it...",
	})

}
