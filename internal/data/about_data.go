package data

import (
	"fmt"

	"github.com/man-on-box/man-on-box.github.io/internal/helpers"
	"github.com/man-on-box/man-on-box.github.io/internal/model"
)

func (d *Data) NewPageAbout() model.PageAbout {
	pageData := newPageData()
	pageData.Head.Title = "Hey, I'm Oli"
	pageData.Head.Desc = "A short and personal piece about my journey into web development."
	pageData.Head.Social = fmt.Sprintf("https://%s/img/social.png", d.SiteDomain)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s/about", d.SiteDomain)

	return model.PageAbout{
		PageData: pageData,
		Content:  helpers.MdFileToHTML("content/about.md", nil).Html,
	}
}
