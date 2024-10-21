package main

import (
	"log"

	"github.com/man-on-box/litepage"
	"github.com/man-on-box/man-on-box.github.io/data"
	"github.com/man-on-box/man-on-box.github.io/pages"
)

func main() {
	domain := "manonbox.io"
	lp, err := litepage.New(domain)
	if err != nil {
		log.Fatalf("Could not create litepage: %v", err)
	}

	data := data.New(domain)

	pages.SetupPages(lp, data)

	lp.BuildOrServe()

}
