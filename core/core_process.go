package core

import (
	"fmt"
	"resource-finder/model"
	"strings"
	"time"
)

func ProcessAllScripts(referenceState []model.ReferenceState, domainList []model.DomainModule) (moveScripts []model.Script) {
	time := createFormattedTime()
	for _, domain := range domainList {
		for _, reference := range referenceState {
			moveScripts = append(moveScripts, processScript(domain, reference, time))
		}
	}
	return moveScripts
}

func createFormattedTime() string {
	time := time.Now()
	return fmt.Sprintf("%d.%d.%d__%d.%d.%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

func processScript(domain model.DomainModule, reference model.ReferenceState, time string) (moveScript model.Script) {
	moveScript.Initialize(domain.Name, reference.Workspace)
	processAllReferenceResources(reference, domain, &moveScript, time)
	return moveScript
}

func processAllReferenceResources(reference model.ReferenceState, domain model.DomainModule, moveScript *model.Script, time string) {
	for referenceSuffix, resources := range reference.Resources {
		proccessReferenceResources(domain, moveScript, referenceSuffix, time, resources)
	}
}

func proccessReferenceResources(domain model.DomainModule, moveScript *model.Script, referenceSuffix, time string, resources []string) {
	if strings.Contains(referenceSuffix, domain.Name) {
		for _, resource := range resources {
			processDomainResources(domain, moveScript, resource, referenceSuffix, time)
			processChildDomain(domain, moveScript, resource, referenceSuffix, time)
		}
	}
}

func processDomainResources(domain model.DomainModule, moveScript *model.Script, resource, referenceSuffix, time string) {
	for _, domainResource := range domain.Resources {
		if strings.Contains(resource, domainResource) {
			moveScript.LoadScriptBody(time, moveScript.Directory, moveScript.Name, resource, referenceSuffix, "", domainResource)
		}
	}
}

func processChildDomain(domain model.DomainModule, moveScript *model.Script, resource, referenceSuffix, time string) {
	for _, childModule := range domain.ChildModules {
		processChildDomainResources(childModule, moveScript, resource, referenceSuffix, time)
	}
}

func processChildDomainResources(childModule model.ChildModule, moveScript *model.Script, resource, referenceSuffix, time string) {
	for _, childModuleResource := range childModule.Resources {
		if strings.Contains(resource, childModuleResource) {
			moveScript.LoadScriptBody(time, moveScript.Directory, moveScript.Name, resource, referenceSuffix, "module."+childModule.Name+".", childModuleResource)
		}
	}
}
