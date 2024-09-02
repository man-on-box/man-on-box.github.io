package data

import (
	"time"
)

const (
	LINKEDIN_URL    = "https://www.linkedin.com/in/oliver-hughes-m880/"
	GITHUB_URL      = "https://github.com/man-on-box"
	GITHUB_REPO_URL = "https://github.com/man-on-box/man-on-box.github.io"
)

var navLinks = []Link{
	{Label: "Home", Url: "/"},
	{Label: "About", Url: "/about.html"},
	{Label: "Articles", Url: "/#articles"},
	{Label: "Skills", Url: "/#skills"},
	{Label: "Experience", Url: "/#experience"},
	{Label: "Contact", Url: "/#contact"},
}

var socials = []Social{
	{Img: "/img/linkedin-logo.svg", Alt: "LinkedIn", Link: Link{Url: LINKEDIN_URL, Label: "LinkedIn"}},
	{Img: "/img/github-logo.svg", Alt: "Github", Link: Link{Url: GITHUB_URL, Label: "Github"}},
}

type PageData struct {
	Head    Head
	NavMenu NavMenu
	Contact Contact
	Footer  Footer
}

type Link struct {
	Url   string
	Label string
}

type Social struct {
	Img  string
	Alt  string
	Link Link
}

type NavMenu struct {
	Links   []Link
	Socials []Social
}

type Contact struct {
	Socials []Social
}

type Head struct {
	Title   string
	Desc    string
	Social  string
	PageUrl string
}

type Footer struct {
	Links       []Link
	Socials     []Social
	RepoUrl     string
	CurrentYear int
}

func newPageData() PageData {
	return PageData{
		Head:    Head{},
		NavMenu: newNavMenu(),
		Contact: newContact(),
		Footer:  newFooter(),
	}
}

func newNavMenu() NavMenu {
	return NavMenu{
		Links:   navLinks,
		Socials: socials,
	}
}

func newContact() Contact {
	return Contact{
		Socials: socials,
	}
}

func newFooter() Footer {
	return Footer{
		Links:       navLinks,
		Socials:     socials,
		RepoUrl:     GITHUB_REPO_URL,
		CurrentYear: time.Now().Year(),
	}
}
