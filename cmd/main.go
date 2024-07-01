package main

import (
	"github.com/man-on-box/man-on-box.github.io/view"
)

func main() {
	static := view.Static{
		DistDir: "dist",
	}

	static.Generate()
}
