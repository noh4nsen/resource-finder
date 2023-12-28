package model

import (
	"regexp"
	"strings"
)

type ReferenceState struct {
	Workspace string
	Resources map[string][]string
}

func (referenceState *ReferenceState) LoadState(workspace string, resources []string) {
	pattern := regexp.MustCompile(`\["([^"]+)"\]`)
	referenceState.Workspace = workspace
	referenceState.Resources = map[string][]string{}
	for _, resource := range resources {
		if strings.Contains(resource, "[") {
			matches := pattern.FindStringSubmatch(resource)
			parts := strings.SplitN(resource, `["`, 2)
			referenceState.Resources[matches[1]] = append(referenceState.Resources[matches[1]], parts[0])
		} else {
			referenceState.Resources["GENERAL"] = append(referenceState.Resources["GENERAL"], resource)
		}
	}
}
