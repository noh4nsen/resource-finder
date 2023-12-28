package core

import (
	"log"
	"resource-finder/cmd"
	"resource-finder/model"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func LoadDomains(config model.Config) (domainResourceList []model.DomainModule) {
	for _, item := range config.Domains {
		var domainResource model.DomainModule
		rootModule := loadModule(item.Path)
		domainResource.Name = item.Domain
		for resource := range rootModule.ManagedResources {
			domainResource.Resources = append(domainResource.Resources, resource)
		}
		log.Printf("\033[0;94m[ %s ]\033[0m :: Loaded resources", domainResource.Name)
		loadChildModules(item.Path, &domainResource, rootModule)
		domainResourceList = append(domainResourceList, domainResource)
	}
	return domainResourceList
}

func loadChildModules(path string, domainResource *model.DomainModule, module *tfconfig.Module) {
	for _, childModule := range module.ModuleCalls {
		var domainResourceChild model.ChildModule
		cmd.Chdir(path)
		module := loadModule(childModule.Source)
		domainResourceChild.Name = childModule.Name
		for resource := range module.ManagedResources {
			domainResourceChild.Resources = append(domainResourceChild.Resources, resource)
		}
		domainResource.ChildModules = append(domainResource.ChildModules, domainResourceChild)
		log.Printf("\033[0;94m[ %s ]\033[0m :: Loaded resources of child module %s", domainResource.Name, domainResourceChild.Name)
	}
}

func loadModule(path string) *tfconfig.Module {
	module, diags := tfconfig.LoadModule(path)
	if diags != nil {
		log.Fatal(diags.Error())
	}
	return module
}
