package data

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/man-on-box/man-on-box.github.io/util"
)

const dateFormat = "02/01/2006"

type article struct {
	Content      template.HTML
	Desc         string
	Published    string
	Title        string
	Slug         string
	Img          string
	ReadTime     string
	SyndicateUrl string
}

type articleMeta struct {
	Title        string `yaml:"title"`
	Slug         string `yaml:"slug"`
	Desc         string `yaml:"description"`
	Img          string `yaml:"img"`
	Published    string `yaml:"published"`
	SyndicateUrl string `yaml:"syndicateUrl"`
}

type pageArticle struct {
	PageData
	Article *article
}

func (a *articleMeta) getTime() time.Time {
	t, err := time.Parse(dateFormat, a.Published)
	if err != nil {
		log.Fatalf("Error parsing date '%s' in article '%s': %v", a.Published, a.Title, err)
	}
	return t
}

func (a *articleMeta) getFormattedDate() string {
	t := a.getTime()
	year, month, day := t.Date()
	return fmt.Sprintf("%d %s, %d", day, month, year)
}

func (d *Data) NewPageArticle(a *article) pageArticle {
	urlPath := fmt.Sprintf("/articles/%s", a.Slug)

	pageData := newPageData()
	pageData.Head.Title = a.Title
	pageData.Head.Desc = a.Desc
	pageData.Head.Social = fmt.Sprintf("https://%s%s", d.SiteDomain, a.Img)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s%s", d.SiteDomain, urlPath)

	return pageArticle{
		PageData: pageData,
		Article:  a,
	}
}

func (d *Data) GetArticles() []article {
	entries, err := os.ReadDir("content/articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	articles := []article{}

	for _, entry := range entries {
		f := "content/articles/" + entry.Name()
		meta := articleMeta{}
		p := util.MdFileToHTML(f, &meta)
		article := article{
			Content:      p.Html,
			Desc:         meta.Desc,
			Published:    meta.getFormattedDate(),
			Title:        meta.Title,
			Slug:         meta.Slug,
			Img:          meta.Img,
			SyndicateUrl: meta.SyndicateUrl,
			ReadTime:     estReadTime(p.WordCount),
		}
		articles = append(articles, article)
	}

	return articles
}

func estReadTime(words int) string {
	wordsPerMinute := 200
	minutes := words / wordsPerMinute
	if minutes < 1 {
		return "1 min read"
	}
	return fmt.Sprintf("%d min read", minutes)
}
