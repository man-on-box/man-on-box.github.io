package data

import (
	"fmt"
	"html/template"

	"github.com/man-on-box/man-on-box.github.io/util"
)

type pageIndex struct {
	PageData
	Content    template.HTML
	Articles   *[]article
	Skills     []skillView
	Experience []job
}

type job struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Date        string   `json:"date"`
	Location    string   `json:"location"`
	Description []string `json:"description"`
}

type skill struct {
	Title string      `json:"title"`
	Color string      `json:"color"`
	Items []skillItem `json:"items"`
}

type skillItem struct {
	Label string `json:"label"`
	Badge string `json:"badge"`
}

type skillView struct {
	Title        string
	BoxClasses   string
	BadgeClasses string
	Items        []skillItem
}

type homeMeta struct {
	Title string `yaml:"title"`
	Desc  string `yaml:"description"`
}

func (d *Data) NewPageIndex(articles *[]article) pageIndex {
	boxClasses := []string{
		"bg-pink-300 shadow-pink-800/60",
		"bg-amber-400 shadow-amber-800/60",
		"bg-cyan-400 shadow-cyan-800/60",
	}

	badgeClasses := []string{
		"text-pink-600",
		"text-amber-600",
		"text-cyan-600",
	}

	pageData := newPageData()
	meta := homeMeta{}
	p := util.MdFileToHTML("content/home/content.md", &meta)

	pageData.Head.Title = meta.Title
	pageData.Head.Desc = meta.Desc
	pageData.Head.Social = fmt.Sprintf("https://%s/img/social.png", d.SiteDomain)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s", d.SiteDomain)

	skills := []skill{}
	skillViews := []skillView{}
	util.ParseJSONFile("content/home/skills.json", &skills)

	for i, skill := range skills {
		skillViews = append(skillViews, skillView{
			Title:        skill.Title,
			BoxClasses:   boxClasses[i%len(boxClasses)],
			BadgeClasses: badgeClasses[i%len(badgeClasses)],
			Items:        skill.Items,
		})
	}

	experience := []job{}
	util.ParseJSONFile("content/home/experience.json", &experience)

	return pageIndex{
		PageData:   pageData,
		Content:    p.Html,
		Articles:   articles,
		Skills:     skillViews,
		Experience: experience,
	}
}
