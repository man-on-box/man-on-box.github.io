package data

import (
	"fmt"
	"html/template"

	"github.com/man-on-box/man-on-box.github.io/util"
)

type PageAbout struct {
	PageData
	Content template.HTML
}

func (d *Data) NewPageAbout() PageAbout {
	pageData := newPageData()
	pageData.Head.Title = "Hey, I'm Oli"
	pageData.Head.Desc = "A short and personal piece about my journey into web development."
	pageData.Head.Social = fmt.Sprintf("https://%s/img/social.png", d.SiteDomain)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s/about", d.SiteDomain)

	return PageAbout{
		PageData: pageData,
		Content:  util.MdFileToHTML("content/about.md", nil).Html,
	}
}
