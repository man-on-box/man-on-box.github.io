package model

import (
	"html/template"
)

type SkillItem struct {
	Label string
	Badge string
}

type Skill struct {
	Title string
	Color string
	Items []SkillItem
}

type Job struct {
	Title       string
	Company     string
	Date        string
	Location    string
	Description []string
}

type PageIndex struct {
	PageData
	Content    template.HTML
	Articles   *[]Article
	Skills     []Skill
	Experience []Job
}
