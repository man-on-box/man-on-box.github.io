// Package static provides a simple way to generate static websites using Go templates and markdown.
package static

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type Static struct {
	DistDir   string
	SiteUrl   string
	tmpl      *template.Template
	sitePaths []string
}

// New creates a new Static instance which can then be used to generate static HTML files, with Markdown support.
func New(distDir string, siteUrl string) *Static {
	s := Static{
		DistDir: distDir,
		SiteUrl: siteUrl,
	}
	s.parseTemplates()
	s.copyPublicDir()
	return &s
}

func (s *Static) Render(tmplName string, filepath string, data interface{}) {
	f := s.createFile(filepath)
	if err := s.tmpl.ExecuteTemplate(f, tmplName, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	s.sitePaths = append(s.sitePaths, filepath)
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

// Done should be called after all pages have been generated.
func (s *Static) Done() {
	s.generateRobotsTxt()
	s.generateSitemap()
}

func (s *Static) generateSitemap() {
	f := s.createFile("/sitemap.xml")
	var builder strings.Builder
	builder.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	builder.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	for _, path := range s.sitePaths {
		urlPath := pathToURL(path)
		builder.WriteString(fmt.Sprintf(`
        <url>
		<loc>https://%s%s</loc>
        </url>`, s.SiteUrl, urlPath))
	}
	builder.WriteString(`</urlset>`)
	f.Write([]byte(builder.String()))
}

func (s *Static) generateRobotsTxt() {
	f := s.createFile("/robots.txt")
	content := `User-agent: *
Disallow:
Allow: /

Sitemap: https://%s/sitemap.xml
`
	robotsTxt := []byte(fmt.Sprintf(content, s.SiteUrl))
	f.Write(robotsTxt)
}

func (s *Static) parseTemplates() {
	path := "./view/templates/**/*.html"
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

func pathToURL(path string) string {
	if path == "/index.html" {
		return "/"
	}
	return strings.TrimSuffix(path, ".html")
}

func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
	}
}
