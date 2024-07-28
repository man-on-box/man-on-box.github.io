package main

import (
	"github.com/man-on-box/man-on-box.github.io/static"
	"github.com/man-on-box/man-on-box.github.io/view"
)

func main() {
	distDir := "dist"
	tmplDir := "./view/templates/**/*.html"

	s := static.New(distDir, tmplDir)
	v := view.New(s)

	v.GeneratePages()
}
