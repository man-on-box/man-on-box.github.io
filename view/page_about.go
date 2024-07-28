package view

import (
	"html/template"
)

type aboutData struct {
	Page    page
	Content template.HTML
}

func (v *View) pageAbout() {
	data := aboutData{
		Page: page{
			Title:   "About Oli",
			Desc:    "Hey I'm Oli, user-centric and product focused software engineer.",
			Nav:     NavLinks,
			Socials: Socials,
		},
		Content: v.static.MdFileToHTML("content/about.md"),
	}

	v.static.Render("about", "/about.html", data)
}
