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
	v.pageIndex()
	v.pageAbout()

	v.static.Done()
}
