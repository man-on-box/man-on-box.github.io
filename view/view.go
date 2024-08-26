// Package view utilizes the static package to generate the static pages and content.
package view

import (
	"github.com/man-on-box/man-on-box.github.io/static"
)

type View struct {
	static *static.Static
}

// New creates a new View instance which can then be used to generate the site as static HTML files.
func New(s *static.Static) *View {
	return &View{s}
}

func (v *View) GeneratePages() {
	defer v.static.Done()
	articles := v.parseArticles()

	v.pageIndex(articles)
	v.pageAbout()

	for _, a := range articles {
		v.pageArticle(a)
	}
}
