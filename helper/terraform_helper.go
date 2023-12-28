package helper

import (
	"strings"
)

func ParseStateList(stdout []byte) (stateList []string) {
	rawStateList := string(stdout)
	stateList = append(stateList, strings.Split(rawStateList, "\n")...)
	return stateList
}

func ParseWorkspaceList(stdout []byte) (workspaceList []string) {
	rawWorkspaceList := string(stdout)
	rawWorkspaceList = strings.Replace(rawWorkspaceList, "default", "", -1)
	rawWorkspaceList = strings.Replace(rawWorkspaceList, " ", "", -1)
	rawWorkspaceList = strings.Replace(rawWorkspaceList, "*", "", -1)

	for _, workspace := range strings.Split(rawWorkspaceList, "\n") {
		if workspace != "" {
			workspaceList = append(workspaceList, workspace)
		}
	}

	return workspaceList
}
