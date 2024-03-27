package model

type DomainModule struct {
	Name         string
	Path         string
	Resources    []string
	ChildModules []ChildModule
}

type ChildModule struct {
	Name      string
	Resources []string
}
