package components

import "time"

type Footer struct {
	Links       []Link
	Socials     []Social
	RepoUrl     string
	CurrentYear int
}

func NewFooter() Footer {
	return Footer{
		Links:       NavLinks,
		Socials:     Socials,
		RepoUrl:     GITHUB_REPO_URL,
		CurrentYear: time.Now().Year(),
	}
}
