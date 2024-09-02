package pages

import (
	"html/template"
	"log"
	"os"

	"github.com/man-on-box/man-on-box.github.io/internal/helpers"
)

func (p *Pages) init() {
	p.tmpl = parseTemplates()
	scaffoldDistDir(p.config.DistDir)
	copyPublicDir(p.config.DistDir)
}

func parseTemplates() *template.Template {
	patterns := []string{
		"./internal/view/*.html",
		"./internal/view/**/*.html",
	}

	tmpl := template.New("")
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
	if err := helpers.CopyDir("public", distDir); err != nil {
		log.Fatalf("Could not copy public directory: %v", err)
	}
}
