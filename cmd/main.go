package main

import (
	"github.com/man-on-box/man-on-box.github.io/pages"
)

func main() {
	pages := pages.New(pages.Config{
		Domain: "manonbox.io",
	})
	pages.Create()
}
