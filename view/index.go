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

type job struct {
	Title       string
	Company     string
	Date        string
	Location    string
	Description []string
}

type pageData struct {
	Title      string
	Skills     []skill
	Experience []job
	Contact    []contact
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
		Experience: []job{
			{
				Title:    "Lead Frontend Developer",
				Company:  "Adevinta",
				Date:     "2019 - Present",
				Location: "Barcelona, Spain",
				Description: []string{
					"Leading the frontend initiative to evolve and build the custom built experimentation platform.",
				},
			}, {
				Title:    "Web Developer",
				Company:  "Discerning Digital",
				Date:     "2017 - 2019",
				Location: "Manchester, UK",
				Description: []string{
					"Working on a variety of projects from small brochure sites to large scale web applications. As well, managing the hosting infrastructure for the company and implementing continuous deployment and testing methodologies.",
				},
			},
			{
				Title:    "Head of Disaster Recovery",
				Company:  "UKFast",
				Date:     "2013 - 2017",
				Location: "Manchester, UK",
				Description: []string{
					"Joined as a Windows Server engineer and quickly moved into the Disaster Recovery team. I was responsible for the design and implementation of the disaster recovery solutions for the company's clients. I also managed the team and the day to day operations of the department.",
				},
			},
		},
	}

	f := s.createFile("/index.html")
	if err := tmpl.ExecuteTemplate(f, "homepage", page); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

}
