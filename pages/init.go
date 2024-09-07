package pages

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/man-on-box/man-on-box.github.io/util"
)

func (p *Pages) init() {
	p.tmpl = parseTemplates()
	scaffoldDistDir(p.config.DistDir)
	copyPublicDir(p.config.DistDir)
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

func scaffoldDistDir(distDir string) {
	os.MkdirAll(distDir+"/articles", os.ModePerm)
}

func copyPublicDir(distDir string) {
	if err := util.CopyDir("public", distDir); err != nil {
		log.Fatalf("Could not copy public directory: %v", err)
	}
}
