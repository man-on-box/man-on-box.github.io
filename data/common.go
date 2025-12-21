package data

const (
	LINKEDIN_URL    = "https://www.linkedin.com/in/oliver-hughes-m880/"
	GITHUB_URL      = "https://github.com/man-on-box"
	GITHUB_REPO_URL = "https://github.com/man-on-box/man-on-box.github.io"
)

var navLinks = []link{
	{Label: "Home", Url: "/"},
	{Label: "About", Url: "/about.html"},
	{Label: "Articles", Url: "/#articles"},
	{Label: "Skills", Url: "/#skills"},
	{Label: "Experience", Url: "/#experience"},
	{Label: "Contact", Url: "/#contact"},
}

var socials = []social{
	{Img: "/img/linkedin-logo.svg", Alt: "LinkedIn", Link: link{Url: LINKEDIN_URL, Label: "LinkedIn"}},
	{Img: "/img/github-logo.svg", Alt: "Github", Link: link{Url: GITHUB_URL, Label: "Github"}},
}

type PageData struct {
	Head    head
	NavMenu navMenu
	Contact contact
	Footer  footer
}

type link struct {
	Url   string
	Label string
}

type social struct {
	Img  string
	Alt  string
	Link link
}

type navMenu struct {
	Links   []link
	Socials []social
}

type contact struct {
	Socials []social
}

type head struct {
	Title   string
	Desc    string
	Social  string
	PageUrl string
}

type footer struct {
	Links   []link
	Socials []social
}

func newPageData() PageData {
	return PageData{
		Head:    head{},
		NavMenu: newNavMenu(),
		Contact: newContact(),
		Footer:  newFooter(),
	}
}

func newNavMenu() navMenu {
	return navMenu{
		Links:   navLinks,
		Socials: socials,
	}
}

func newContact() contact {
	return contact{
		Socials: socials,
	}
}

func newFooter() footer {
	return footer{
		Links:   navLinks,
		Socials: socials,
	}
}
