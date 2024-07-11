package view

const (
	LINKEDIN_URL = "https://www.linkedin.com/in/oliver-hughes-m880/"
	GITHUB_URL   = "https://github.com/man-on-box"
)

type Link struct {
	Url   string
	Label string
}

type page struct {
	Title   string
	Desc    string
	Contact []Link
	Nav     []Link
}

var Contact = []Link{
	{Label: "LinkedIn", Url: LINKEDIN_URL},
	{Label: "Github", Url: GITHUB_URL},
}

var NavLinks = []Link{
	{Label: "Home", Url: "/"},
	// {Label: "About", Url: "/about.html"},
	{Label: "Skills", Url: "/#skills"},
	{Label: "Experience", Url: "/#experience"},
	{Label: "Contact", Url: "/#contact"},
}
