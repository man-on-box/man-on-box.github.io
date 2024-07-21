package view

import (
	"html/template"
)

type aboutData struct {
	Page    page
	Content template.HTML
}

func (s *Static) pageAbout() {
	data := aboutData{
		Page: page{
			Title:   "About Oli",
			Desc:    "Hey I'm Oli, user-centric and product focused software engineer.",
			Nav:     NavLinks,
			Socials: Socials,
		},
		Content: mdFileToHTML("content/about.md"),
	}

	s.render("about", "/about.html", data)
}
