package pages

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"

	"github.com/man-on-box/litepage"
	"github.com/man-on-box/man-on-box.github.io/data"
)

func SetupPages(lp litepage.Litepage, data *data.Data) {
	render := parseTemplates()
	articles := data.GetArticles()

	lp.Page("/index.html", func(w io.Writer) {
		d := data.NewPageIndex(&articles)
		render(w, "page-index", d)
	})

	lp.Page("/about.html", func(w io.Writer) {
		d := data.NewPageAbout()
		render(w, "page-about", d)
	})

	for _, a := range articles {
		filePath := fmt.Sprintf("/articles/%s.html", a.Slug)
		lp.Page(filePath, func(w io.Writer) {
			d := data.NewPageArticle(&a)
			render(w, "page-article", d)
		})
	}
}

func parseTemplates() (render func(w io.Writer, name string, data interface{})) {
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

	return func(w io.Writer, name string, data interface{}) {
		err := tmpl.ExecuteTemplate(w, name, data)
		if err != nil {
			log.Fatalf("Error executing template '%s': %v", name, err)
		}
	}
}
