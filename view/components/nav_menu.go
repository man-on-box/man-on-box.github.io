package components

type NavMenu struct {
	Links   []Link
	Socials []Social
}

func NewNavMenu() NavMenu {
	return NavMenu{
		Links:   NavLinks,
		Socials: Socials,
	}
}
