package pages

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	"github.com/man-on-box/man-on-box.github.io/internal/data"
	"github.com/man-on-box/man-on-box.github.io/internal/util"
)

type Page struct {
	render   func(w io.Writer)
	filePath string
}

type Config struct {
	DistDir string
	Domain  string
}

type Pages struct {
	config Config
	data   *data.Data
	tmpl   *template.Template
	pages  *[]Page
}

func New(config Config) *Pages {
	p := Pages{
		config: config,
		data:   data.New(config.Domain),
	}
	return &p
}

func (p *Pages) Build() {
	fmt.Println("Building static site: " + p.config.Domain)
	p.init()
	p.pages = p.setupPages()
	p.createPages()
	p.postBuild()
}

func (p *Pages) postBuild() {
	p.generateRobotsTxt()
	p.generateSitemap()
}

func (p *Pages) executeTemplate(w io.Writer, name string, data any) {
	err := p.tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}

func (p *Pages) createPages() {
	for _, page := range *p.pages {
		fmt.Println("Creating page: ", page.filePath)
		f := util.CreateFile(p.config.DistDir + page.filePath)
		page.render(f)
	}

}

func (p *Pages) generateRobotsTxt() {
	f := util.CreateFile(p.config.DistDir + "/robots.txt")
	content := `User-agent: *
Disallow:
Allow: /

Sitemap: https://%s/sitemap.xml
`
	robotsTxt := []byte(fmt.Sprintf(content, p.config.Domain))
	f.Write(robotsTxt)
}

func (p *Pages) generateSitemap() {
	f := util.CreateFile(p.config.DistDir + "/sitemap.xml")
	var builder strings.Builder
	builder.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	builder.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	for _, page := range *p.pages {
		urlPath := fileToUrlPath(page.filePath)
		builder.WriteString(fmt.Sprintf(`
        <url>
		<loc>https://%s%s</loc>
        </url>`, p.config.Domain, urlPath))
	}
	builder.WriteString(`</urlset>`)
	f.Write([]byte(builder.String()))
}

func fileToUrlPath(path string) string {
	if path == "/index.html" {
		return "/"
	}
	return strings.TrimSuffix(path, ".html")
}