package view

import (
	"html/template"
	"log"
)

type aboutData struct {
	Page page
}

func (s *Static) pageAbout(tmpl *template.Template) {
	data := aboutData{
		Page: page{
			Title:   "Hey, I'm Oli",
			Nav:     NavLinks,
			Socials: Socials,
			Desc:    "Hey I'm Oli, user-centric and product focused software engineer.",
		},
	}

	f := s.createFile("/about.html")
	if err := tmpl.ExecuteTemplate(f, "about", data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
