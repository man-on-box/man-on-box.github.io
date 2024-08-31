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
	urlPath := fmt.Sprintf("/articles/%s", a.Slug)
	filePath := fmt.Sprintf("%s.html", urlPath)
	data := articleData{
		Head: components.Head{
			Title:   a.Title,
			Desc:    a.Desc,
			Social:  fmt.Sprintf("https://%s%s", v.static.SiteUrl, a.Img),
			PageUrl: fmt.Sprintf("https://%s%s", v.static.SiteUrl, urlPath),
		},
		NavMenu: components.NewNavMenu(),
		Contact: components.NewContact(),
		Footer:  components.NewFooter(),
		Article: a,
	}

	v.static.Render("page-article", filePath, data)
}
