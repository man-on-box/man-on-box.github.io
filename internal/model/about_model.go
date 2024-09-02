package model

import (
	"html/template"
)

type PageAbout struct {
	PageData
	Content template.HTML
}
