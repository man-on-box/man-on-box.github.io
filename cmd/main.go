package main

import (
	"github.com/man-on-box/man-on-box.github.io/static"
	"github.com/man-on-box/man-on-box.github.io/view"
)

func main() {
	distDir := "dist"
	siteUrl := "manonbox.io"

	s := static.New(distDir, siteUrl)
	v := view.New(s)

	v.GeneratePages()
}
