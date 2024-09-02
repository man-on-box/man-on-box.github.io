package data

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/man-on-box/man-on-box.github.io/internal/helpers"
)

const dateFormat = "02/01/2006"

type Article struct {
	Content   template.HTML
	Desc      string
	Published string
	Title     string
	Slug      string
	Img       string
	ReadTime  string
}

type ArticleMeta struct {
	Title     string `yaml:"title"`
	Slug      string `yaml:"slug"`
	Desc      string `yaml:"description"`
	Img       string `yaml:"img"`
	Published string `yaml:"published"`
}

type PageArticle struct {
	PageData
	Article *Article
}

func (a *ArticleMeta) getTime() time.Time {
	t, err := time.Parse(dateFormat, a.Published)
	if err != nil {
		log.Fatalf("Error parsing date '%s' in article '%s': %v", a.Published, a.Title, err)
	}
	return t
}

func (a *ArticleMeta) GetFormattedDate() string {
	t := a.getTime()
	year, month, day := t.Date()
	return fmt.Sprintf("%d %s, %d", day, month, year)
}

func (d *Data) NewPageArticle(a *Article) PageArticle {
	urlPath := fmt.Sprintf("/articles/%s", a.Slug)

	pageData := newPageData()
	pageData.Head.Title = a.Title
	pageData.Head.Desc = a.Desc
	pageData.Head.Social = fmt.Sprintf("https://%s%s", d.SiteDomain, a.Img)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s%s", d.SiteDomain, urlPath)

	return PageArticle{
		PageData: pageData,
		Article:  a,
	}
}

func (d *Data) GetArticles() []Article {
	entries, err := os.ReadDir("content/articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	articles := []Article{}

	for _, entry := range entries {
		f := "content/articles/" + entry.Name()
		meta := ArticleMeta{}
		p := helpers.MdFileToHTML(f, &meta)
		article := Article{
			Content:   p.Html,
			Desc:      meta.Desc,
			Published: meta.GetFormattedDate(),
			Title:     meta.Title,
			Slug:      meta.Slug,
			Img:       meta.Img,
			ReadTime:  estReadTime(p.WordCount),
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
