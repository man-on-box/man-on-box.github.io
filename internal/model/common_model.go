package model

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
