package model

import (
	"fmt"
	"html/template"
	"log"
	"time"
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
