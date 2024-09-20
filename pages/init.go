package pages

import (
	"flag"
	"html/template"
	"io"
	"log"
	"time"

	"github.com/man-on-box/man-on-box.github.io/data"
)

const distDir = "dist"

type Page struct {
	render   func(w io.Writer)
	filePath string
}

type Config struct {
	Domain string
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

	p.tmpl = parseTemplates()
	p.pages = p.setupPages()
	return &p
}

func (p *Pages) Create() {
	dev := flag.Bool("dev", false, "Serve pages via dev server")
	port := flag.String("port", "3000", "Server port to start on")
	flag.Parse()
	if *dev {
		p.serve(*port)
	} else {
		p.build()
	}
}

func parseTemplates() *template.Template {
	patterns := []string{
		"./view/*.html",
		"./view/**/*.html",
	}

	tmpl := template.New("").Funcs(template.FuncMap{
		"version": func() string {
			return time.Now().Format("01021504")
		},
	})
	var err error

	for _, pattern := range patterns {
		tmpl, err = tmpl.ParseGlob(pattern)
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
	}

	return tmpl
}

func (p *Pages) executeTemplate(w io.Writer, name string, data any) {
	err := p.tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}

/*
func (p *Pages) init() {
	p.tmpl = parseTemplates()
	scaffoldDistDir(p.config.DistDir)
	copyPublicDir(p.config.DistDir)
}

func parseTemplates() *template.Template {
	patterns := []string{
		"./view/*.html",
	}

	tmpl := template.New("").Funcs(template.FuncMap{
		"version": func() string {
			return time.Now().Format("01021504")
		},
	})
	var err error

	for _, pattern := range patterns {
		tmpl, err = tmpl.ParseGlob(pattern)
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
	}

	return tmpl
}

func scaffoldDistDir(distDir string) {
	os.MkdirAll(distDir+"/articles", os.ModePerm)
}

func copyPublicDir(distDir string) {
	if err := util.CopyDir("public", distDir); err != nil {
		log.Fatalf("Could not copy public directory: %v", err)
	}
}
*/
