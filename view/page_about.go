package view

import (
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
			Title:  "About Oli",
			Desc:   "A short and personal piece about my journey into Web Development.",
			Social: v.data.socialImgUrl,
		},
		NavMenu: components.NewNavMenu(),
		Contact: components.NewContact(),
		Footer:  components.NewFooter(),
		Content: v.static.MdFileToHTML("content/about.md", nil).Html,
	}

	v.static.Render("page-about", "/about.html", data)
}
