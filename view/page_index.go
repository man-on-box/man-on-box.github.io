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

type job struct {
	Title       string
	Company     string
	Date        string
	Location    string
	Description []string
}

type indexData struct {
	Page       page
	Skills     []skill
	Experience []job
}

func (s *Static) pageIndex(tmpl *template.Template) {

	data := indexData{
		Page: page{
			Title:   "Hey, I'm Oli",
			Nav:     NavLinks,
			Socials: Socials,
			Desc:    "Hey I'm Oli, user-centric and product focused software engineer.",
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
					{Label: "HTMX", Badge: "</>"},
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
					{Label: "Git", Badge: "{ }"},
					{Label: "Amazon Web Services", Badge: "AWS"},
					{Label: "Linux Administation", Badge: "/"},
					{Label: "Windows Server", Badge: "C:\\"},
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
					"Leading the frontend initiative to evolve and build the in-house built experimentation platform and libraries used across European marketplaces with millions of users.",
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
			{Title: "Head of Disaster Recovery",
				Company:  "UKFast",
				Date:     "2015 - 2017",
				Location: "Manchester, UK",
				Description: []string{
					"Moving from the Senior Windows team to take lead of the Disaster Recovery team, I was responsible for the design and implementation of the disaster recovery solutions for the company's enterprise clients, as well as architecting and scaling the backup infrastructure to handle petabytes of data distibuted across multiple datacenters. I also managed the team and the day to day operations of the department.",
				},
			}, {
				Title:    "Windows Server Engineer",
				Company:  "UKFast",
				Date:     "2013 - 2015",
				Location: "Manchester, UK",
				Description: []string{
					"Starting in the company as a Windows Engineer, day to day duties comprised of supporting customers via phone or ticket system for their solutions hosted with UK Fast. This included but not limited to support with Windows 2008/2012, Active Directory, Exchange, IIS, VMWare, Hyper-V, MSSQL and basic support on CentOS, Red Hat and Ubuntu servers.",
					"Within a few months I was then promoted to a Senior Windows Engineer, which involved supporting customers of a higher tier with more complex solutions such as clusters, DAGs, MSSQL High Availability and multi node Exchange environments.",
				},
			}, {
				Title:    "Technical Support Analyst",
				Company:  "NorthgateArinso",
				Date:     "2008 - 2013",
				Location: "Manchester, UK",
				Description: []string{
					"Providing 1st to 3rd line technical support to a user base of 500+. Daily activities ranged widely from repairing laptops, software troubleshooting remotely or face to face, to server administration, managing Active Directory and Exchange servers.",
				},
			},
		},
	}

	f := s.createFile("/index.html")
	if err := tmpl.ExecuteTemplate(f, "homepage", data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

}
