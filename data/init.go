package data

type Data struct {
	SiteDomain string
}

func New(domain string) *Data {
	return &Data{
		SiteDomain: domain,
	}
}
