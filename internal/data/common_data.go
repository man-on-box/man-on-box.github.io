package data

import (
	"time"

	"github.com/man-on-box/man-on-box.github.io/internal/model"
)

const (
	LINKEDIN_URL    = "https://www.linkedin.com/in/oliver-hughes-m880/"
	GITHUB_URL      = "https://github.com/man-on-box"
	GITHUB_REPO_URL = "https://github.com/man-on-box/man-on-box.github.io"
)

var navLinks = []model.Link{
	{Label: "Home", Url: "/"},
	{Label: "About", Url: "/about.html"},
	{Label: "Articles", Url: "/#articles"},
	{Label: "Skills", Url: "/#skills"},
	{Label: "Experience", Url: "/#experience"},
	{Label: "Contact", Url: "/#contact"},
}

var socials = []model.Social{
	{Img: "/img/linkedin-logo.svg", Alt: "LinkedIn", Link: model.Link{Url: LINKEDIN_URL, Label: "LinkedIn"}},
	{Img: "/img/github-logo.svg", Alt: "Github", Link: model.Link{Url: GITHUB_URL, Label: "Github"}},
}

func newPageData() model.PageData {
	return model.PageData{
		Head:    model.Head{},
		NavMenu: newNavMenu(),
		Contact: newContact(),
		Footer:  newFooter(),
	}
}

func newNavMenu() model.NavMenu {
	return model.NavMenu{
		Links:   navLinks,
		Socials: socials,
	}
}

func newContact() model.Contact {
	return model.Contact{
		Socials: socials,
	}
}

func newFooter() model.Footer {
	return model.Footer{
		Links:       navLinks,
		Socials:     socials,
		RepoUrl:     GITHUB_REPO_URL,
		CurrentYear: time.Now().Year(),
	}
}
