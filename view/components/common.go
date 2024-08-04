package components

const (
	LINKEDIN_URL    = "https://www.linkedin.com/in/oliver-hughes-m880/"
	GITHUB_URL      = "https://github.com/man-on-box"
	GITHUB_REPO_URL = "https://github.com/man-on-box/man-on-box.github.io"
)

type Link struct {
	Url   string
	Label string
}

type Social struct {
	Img  string
	Alt  string
	Link Link
}

type page struct {
	Title   string
	Desc    string
	Nav     []Link
	Socials []Social
}

var NavLinks = []Link{
	{Label: "Home", Url: "/"},
	{Label: "About", Url: "/about.html"},
	{Label: "Skills", Url: "/#skills"},
	{Label: "Experience", Url: "/#experience"},
	{Label: "Contact", Url: "/#contact"},
}

var Socials = []Social{
	{Img: "/img/linkedin-logo.svg", Alt: "LinkedIn", Link: Link{Url: LINKEDIN_URL, Label: "LinkedIn"}},
	{Img: "/img/github-logo.svg", Alt: "Github", Link: Link{Url: GITHUB_URL, Label: "Github"}},
}
