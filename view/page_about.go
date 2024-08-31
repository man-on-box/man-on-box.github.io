package view

import (
	"fmt"
	"html/template"

	"github.com/man-on-box/man-on-box.github.io/view/components"
)

type aboutData struct {
	Head    components.Head
	NavMenu components.NavMenu
	Contact components.Contact
	Footer  components.Footer
	Content template.HTML
}

func (v *View) pageAbout() {
	data := aboutData{
		Head: components.Head{
			Title:   "About Oli",
			Desc:    "A short and personal piece about my journey into web development.",
			Social:  v.data.socialImgUrl,
			PageUrl: fmt.Sprintf("https://%s/about", v.static.SiteUrl),
		},
		NavMenu: components.NewNavMenu(),
		Contact: components.NewContact(),
		Footer:  components.NewFooter(),
		Content: v.static.MdFileToHTML("content/about.md", nil).Html,
	}

	v.static.Render("page-about", "/about.html", data)
}
