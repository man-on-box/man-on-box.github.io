package main

import (
	"github.com/man-on-box/man-on-box.github.io/internal/pages"
)

func main() {
	// litepage.New
	pages := pages.New(pages.Config{
		DistDir: "dist",
		Domain:  "manonbox.io",
	})
	pages.Build()
}
