package components

type Contact struct {
	Socials []Social
}

func NewContact() Contact {
	return Contact{
		Socials: Socials,
	}
}
