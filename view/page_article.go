package view

import (
	"fmt"

	"github.com/man-on-box/man-on-box.github.io/view/components"
)

type articleData struct {
	Head    components.Head
	NavMenu components.NavMenu
	Contact components.Contact
	Footer  components.Footer
	Article article
}

func (v *View) pageArticle(a article) {
	data := articleData{
		Head: components.Head{
			Title: a.Title,
			Desc:  a.Desc,
		},
		NavMenu: components.NewNavMenu(),
		Contact: components.NewContact(),
		Footer:  components.NewFooter(),
		Article: a,
	}

	f := fmt.Sprintf("/articles/%s.html", a.Slug)
	v.static.Render("page-article", f, data)
}
