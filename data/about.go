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

type aboutMeta struct {
	Title string `yaml:"title"`
	Desc  string `yaml:"description"`
}

func (d *Data) NewPageAbout() PageAbout {
	meta := aboutMeta{}
	p := util.MdFileToHTML("content/about/content.md", &meta)

	pageData := newPageData()
	pageData.Head.Title = meta.Title
	pageData.Head.Desc = meta.Desc
	pageData.Head.Social = fmt.Sprintf("https://%s/img/social.png", d.SiteDomain)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s/about", d.SiteDomain)

	return PageAbout{
		PageData: pageData,
		Content:  p.Html,
	}
}
