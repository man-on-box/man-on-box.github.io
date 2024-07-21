package view

import (
	"html/template"
	"log"
)

type aboutData struct {
	Page    page
	Content template.HTML
}

func (s *Static) pageAbout(tmpl *template.Template) {
	data := aboutData{
		Page: page{
			Title:   "Hey, I'm Oli",
			Nav:     NavLinks,
			Socials: Socials,
			Desc:    "Hey I'm Oli, user-centric and product focused software engineer.",
		},
		Content: mdFileToHTML("content/about.md"),
	}

	f := s.createFile("/about.html")
	if err := tmpl.ExecuteTemplate(f, "about", data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
