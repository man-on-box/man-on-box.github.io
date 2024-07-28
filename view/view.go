package view

import (
	"github.com/man-on-box/man-on-box.github.io/static"
)

type View struct {
	static *static.Static
}

func New(s *static.Static) *View {
	return &View{s}
}

func (v *View) GeneratePages() {
	v.pageIndex()
	v.pageAbout()
}
