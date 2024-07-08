package view

import (
	"html/template"
	"log"
)

type item struct {
	Label string
	Badge string
}

type skill struct {
	Title string
	Color string
	Items []item
}

type contact struct {
	Link  string
	Label string
}

type pageData struct {
	Title   string
	Skills  []skill
	Contact []contact
}

func (s *Static) pageIndex(tmpl *template.Template) {

	page := pageData{
		Title: "Man on Box",
		Contact: []contact{
			{Label: "LinkedIn", Link: "https://www.linkedin.com/in/oliver-hughes-m880/"},
			{Label: "Github", Link: "https://github.com/man-on-box"},
		},
		Skills: []skill{
			{
				Title: "Frontend",
				Color: "pink",
				Items: []item{
					{Label: "Javascript", Badge: "JS"},
					{Label: "Typescript", Badge: "TS"},
					{Label: "React", Badge: "</>"},
					{Label: "Astro", Badge: "</>"},
					{Label: "Vite", Badge: "{ }"},
					{Label: "Webpack", Badge: "{ }"},
					{Label: "Tailwind", Badge: "{ }"},
					{Label: "SASS", Badge: "{ }"},
					{Label: "Jest", Badge: "{ }"},
					{Label: "Cypress", Badge: "{ }"},
					{Label: "React Testing Library", Badge: "RTL"},
				},
			},
			{
				Title: "Backend",
				Color: "cyan",
				Items: []item{
					{Label: "Golang", Badge: "GO"},
					{Label: "NodeJs", Badge: "{ }"},
					{Label: "Express", Badge: "{ }"},
					{Label: "Postgres", Badge: "{ }"},
					{Label: "Redis", Badge: "{ }"},
					{Label: "oAuth", Badge: "{ }"},
				},
			},
			{
				Title: "DevOps",
				Color: "amber",
				Items: []item{
					{Label: "Amazon Web Services", Badge: "AWS"},
					{Label: "Github Actions", Badge: "GHA"},
					{Label: "Travis", Badge: "{ }"},
					{Label: "Docker", Badge: "{ }"},
					{Label: "Kubernetes", Badge: "{ }"},
				},
			},
		},
	}

	f := s.createFile("/index.html")
	if err := tmpl.ExecuteTemplate(f, "homepage", page); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

}
