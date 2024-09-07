package pages

import (
	"fmt"
	"io"
)

func (p *Pages) setupPages() *[]Page {
	articles := p.data.GetArticles()

	pages := []Page{
		{
			filePath: "/index.html",
			render: func(w io.Writer) {
				d := p.data.NewPageIndex(&articles)
				p.executeTemplate(w, "page-index", d)
			},
		},
		{
			filePath: "/about.html",
			render: func(w io.Writer) {
				d := p.data.NewPageAbout()
				p.executeTemplate(w, "page-about", d)
			},
		},
	}

	for _, a := range articles {
		pages = append(pages, Page{
			filePath: fmt.Sprintf("/articles/%s.html", a.Slug),
			render: func(w io.Writer) {
				d := p.data.NewPageArticle(&a)
				p.executeTemplate(w, "page-article", d)
			},
		})
	}

	return &pages
}
