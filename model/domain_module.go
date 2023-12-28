package model

type DomainModule struct {
	Name         string
	Resources    []string
	ChildModules []ChildModule
}

type ChildModule struct {
	Name      string
	Resources []string
}
