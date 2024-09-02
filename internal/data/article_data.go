package data

import (
	"fmt"
	"log"
	"os"

	"github.com/man-on-box/man-on-box.github.io/internal/helpers"
	"github.com/man-on-box/man-on-box.github.io/internal/model"
)

func (d *Data) NewPageArticle(a *model.Article) model.PageArticle {
	urlPath := fmt.Sprintf("/articles/%s", a.Slug)

	pageData := newPageData()
	pageData.Head.Title = a.Title
	pageData.Head.Desc = a.Desc
	pageData.Head.Social = fmt.Sprintf("https://%s%s", d.SiteDomain, a.Img)
	pageData.Head.PageUrl = fmt.Sprintf("https://%s%s", d.SiteDomain, urlPath)

	return model.PageArticle{
		PageData: pageData,
		Article:  a,
	}
}

func (d *Data) GetArticles() []model.Article {
	entries, err := os.ReadDir("content/articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	articles := []model.Article{}

	for _, entry := range entries {
		f := "content/articles/" + entry.Name()
		meta := model.ArticleMeta{}
		p := helpers.MdFileToHTML(f, &meta)
		article := model.Article{
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
