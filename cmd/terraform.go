package cmd

import (
	"log"
	"os/exec"
	"resource-finder/helper"
)

var TERRAFORM string = "terraform"

func StateList(path, workspace string) (stateList []string) {
	Chdir(path)
	Which(TERRAFORM)
	workspaceSelect(workspace)
	cmd := exec.Command(TERRAFORM, "state", "list")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalf("\033[0;91m[ TERRAFORM ]\033[0m :: Could not execute `terraform state list`: %s", err.Error())
	}
	log.Printf("\033[0;95m[ TERRAFORM ]\033[0m :: Executed `terraform state list` in %s", path)
	return helper.ParseStateList(stdout)
}

func WorkspaceList(path string) (workspaceList []string) {
	Chdir(path)
	Which(TERRAFORM)
	cmd := exec.Command(TERRAFORM, "workspace", "list")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalf("\033[0;91m[ TERRAFORM ]\033[0m :: Could not execute `terraform workspace list`: %s", err.Error())
	}
	workspaceList = helper.ParseWorkspaceList(stdout)
	log.Printf("\033[0;95m[ TERRAFORM ]\033[0m :: Executed `terraform workspace select`: %s", workspaceList)

	return workspaceList
}

func workspaceSelect(workspace string) {
	cmd := exec.Command(TERRAFORM, "workspace", "select", workspace)
	_, err := cmd.Output()
	if err != nil {
		log.Fatalf("\033[0;91m TERRAFORM ]\033[0m :: Could not execute `terraform workspace select`: %s", err.Error())
	}
	log.Printf("\033[0;95m[ TERRAFORM ]\033[0m :: Executed `terraform workspace select %s`", workspace)
}
