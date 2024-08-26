package view

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

type article struct {
	Content   template.HTML
	Desc      string
	Published string
	Title     string
	Slug      string
	Img       string
}

type articleMeta struct {
	Title     string `yaml:"title"`
	Slug      string `yaml:"slug"`
	Desc      string `yaml:"description"`
	Img       string `yaml:"img"`
	Published string `yaml:"published"`
}

var dateFormat = "02/01/2006"

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

func (v *View) parseArticles() []article {
	entries, err := os.ReadDir("content/articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	articles := []article{}

	for _, entry := range entries {
		log.Println(entry.Name())
		f := "content/articles/" + entry.Name()
		meta := articleMeta{}
		content := v.static.MdFileToHTML(f, &meta)
		article := article{
			Content:   content,
			Desc:      meta.Desc,
			Published: meta.getFormattedDate(),
			Title:     meta.Title,
			Slug:      meta.Slug,
			Img:       meta.Img,
		}
		articles = append(articles, article)
	}

	return articles
}
