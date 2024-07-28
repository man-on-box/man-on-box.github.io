// Package static provides a simple way to generate static websites using Go templates and markdown.
package static

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/yuin/goldmark"
)

type Static struct {
	DistDir string
	tmpl    *template.Template
}

// New creates a new Static instance which can then be used to generate static HTML files, with Markdown support.
func New(distDir string, tmplDir string) *Static {
	s := Static{
		DistDir: distDir,
	}
	s.parseTemplates(tmplDir)
	s.copyPublicDir()
	return &s
}

func (s *Static) Render(tmplName string, filepath string, data interface{}) {
	f := s.createFile(filepath)
	if err := s.tmpl.ExecuteTemplate(f, tmplName, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func (s *Static) MdFileToHTML(file string) template.HTML {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	if err := goldmark.Convert(f, &buf); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	return template.HTML(buf.String())
}

func (s *Static) MdToHTML(md string) template.HTML {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	return template.HTML(buf.String())
}

func (s *Static) parseTemplates(path string) {
	tmpl, err := template.New("").Funcs(createFuncMap()).ParseGlob(path)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	s.tmpl = tmpl
}

func (s *Static) copyPublicDir() {
	if err := copyDir("public", s.DistDir); err != nil {
		log.Fatalf("Could not copy public directory: %v", err)
	}
}

func (s *Static) createFile(fileName string) *os.File {
	f, err := os.Create(s.DistDir + fileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	return f
}

func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
	}
}
