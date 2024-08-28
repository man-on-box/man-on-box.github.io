// Package view utilizes the static package to generate the static pages and content.
package view

import (
	"fmt"

	"github.com/man-on-box/man-on-box.github.io/static"
)

type viewData struct {
	socialImgUrl string
}

type View struct {
	static *static.Static
	data   viewData
}

// New creates a new View instance which can then be used to generate the site as static HTML files.
func New(s *static.Static) *View {
	data := viewData{
		socialImgUrl: fmt.Sprintf("https://%s/img/social.png", s.SiteUrl),
	}
	return &View{
		static: s,
		data:   data,
	}
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
