package components

import "time"

type Footer struct {
	Links       []Link
	Socials     []Social
	GithubUrl   string
	CurrentYear int
}

func NewFooter() Footer {
	return Footer{
		Links:       NavLinks,
		Socials:     Socials,
		GithubUrl:   "https://github.com/man-on-box/man-on-box.github.io",
		CurrentYear: time.Now().Year(),
	}
}
