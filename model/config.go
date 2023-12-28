package model

type Config struct {
	Reference Reference
	Domains   []Domain
}

type Reference struct {
	Path string
}

type Domain struct {
	Domain string
	Path   string
}
