package data

import (
	"fmt"
	"html/template"

	"github.com/man-on-box/man-on-box.github.io/util"
)

type PageIndex struct {
	PageData
	Content    template.HTML
	Articles   *[]Article
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

var boxClassMap = map[string]string{
	"pink":  "border-pink-800 hover:border-pink-500 shadow-pink-800/60 hover:shadow-pink-800",
	"amber": "border-amber-800 hover:border-amber-500 shadow-amber-800/60 hover:shadow-amber-800",
	"cyan":  "border-cyan-800 hover:border-cyan-500 shadow-cyan-800/60 hover:shadow-cyan-800",
}

var badgeClassMap = map[string]string{
	"pink":  "text-pink-600",
	"amber": "text-amber-600",
	"cyan":  "text-cyan-600",
}

func (d *Data) NewPageIndex(articles *[]Article) PageIndex {
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

	for _, skill := range skills {
		skillViews = append(skillViews, skillView{
			Title:        skill.Title,
			BoxClasses:   boxClassMap[skill.Color],
			BadgeClasses: badgeClassMap[skill.Color],
			Items:        skill.Items,
		})
	}

	experience := []job{}
	util.ParseJSONFile("content/home/experience.json", &experience)

	return PageIndex{
		PageData:   pageData,
		Content:    p.Html,
		Articles:   articles,
		Skills:     skillViews,
		Experience: experience,
	}
}
