package view

import (
	"html/template"

	"github.com/man-on-box/man-on-box.github.io/view/components"
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

type job struct {
	Title       string
	Company     string
	Date        string
	Location    string
	Description []string
}

type pageData struct {
	Head       components.Head
	NavMenu    components.NavMenu
	Contact    components.Contact
	Footer     components.Footer
	Content    template.HTML
	Articles   []article
	Skills     []skill
	Experience []job
}

func (v *View) pageIndex(articles []article) {

	data := pageData{
		Head: components.Head{
			Title: "Hey, I'm Oli",
			Desc:  "Hey I'm Oli, user-centric and product focused software engineer.",
		},
		NavMenu:  components.NewNavMenu(),
		Contact:  components.NewContact(),
		Footer:   components.NewFooter(),
		Content:  v.static.MdFileToHTML("content/home.md", nil),
		Articles: articles,
		Skills: []skill{
			{
				Title: "Frontend",
				Color: "pink",
				Items: []item{
					{Label: "Javascript", Badge: "JS"},
					{Label: "Typescript", Badge: "TS"},
					{Label: "React", Badge: "</>"},
					{Label: "Astro", Badge: "</>"},
					{Label: "HTMX", Badge: "</>"},
					{Label: "Vite", Badge: "{ }"},
					{Label: "Webpack", Badge: "{ }"},
					{Label: "Tailwind", Badge: "{ }"},
					{Label: "HTML", Badge: "</>"},
					{Label: "CSS", Badge: "{ }"},
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
					{Label: "Linux Administation", Badge: "$_"},
					{Label: "Windows Server", Badge: "C:\\"},
					{Label: "Git", Badge: "{ }"},
					{Label: "Github Actions", Badge: "GHA"},
					{Label: "Travis", Badge: "{ }"},
					{Label: "Docker", Badge: "{ }"},
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
					"I am currently leading the frontend initiative to evolve and build the in-house built experimentation platform and libraries used across European marketplaces with millions of users.",
					"In addition, I am working across a large tech stack with a focus on React, Typescript, Vite, NodeJS and Golang, and helping and mentoring fellow developers.",
				},
			}, {
				Title:    "Web Developer",
				Company:  "Discerning Digital",
				Date:     "2017 - 2019",
				Location: "Manchester, UK",
				Description: []string{
					"My job involved working on a variety of projects from small brochure sites to large scale web applications. Furthermore, I managed the hosting infrastructure for the company and implemented continuous deployment and testing methodologies.",
				},
			},
			{Title: "Head of Disaster Recovery",
				Company:  "UKFast",
				Date:     "2015 - 2017",
				Location: "Manchester, UK",
				Description: []string{
					"Moving from the Senior Windows team to take lead of the Disaster Recovery team, I was responsible for the design and implementation of the disaster recovery solutions for the company's enterprise clients, as well as architecting and scaling the backup infrastructure to handle petabytes of data distributed across multiple datacenters. I also managed the team and the day to day operations of the department.",
				},
			}, {
				Title:    "Windows Server Engineer",
				Company:  "UKFast",
				Date:     "2013 - 2015",
				Location: "Manchester, UK",
				Description: []string{
					"Starting in the company as a Windows Engineer, day to day duties consisted of supporting customers via phone or ticket system for their solutions hosted with UK Fast. This included (but was not limited to) support with Windows 2008/2012, Active Directory, Exchange, IIS, VMWare, Hyper-V, MSSQL and basic support on CentOS, Red Hat and Ubuntu servers.",
					"Within a few months I was then promoted to a Senior Windows Engineer, which involved supporting customers of a higher tier with more complex solutions such as clusters, DAGs, MSSQL High Availability and multi node Exchange environments.",
				},
			}, {
				Title:    "Technical Support Analyst",
				Company:  "NorthgateArinso",
				Date:     "2008 - 2013",
				Location: "Manchester, UK",
				Description: []string{
					"I provided 1st to 3rd line technical support to a user base of 500+. Daily activities ranged widely from repairing laptops, software troubleshooting remotely or face to face, to server administration, managing Active Directory and Exchange servers.",
				},
			},
		},
	}

	v.static.Render("page-home", "/index.html", data)
}
