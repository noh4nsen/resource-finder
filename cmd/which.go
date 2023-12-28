package cmd

import (
	"log"
	"os/exec"
)

func Which(application string) {
	cmd := exec.Command("which", application)
	_, err := cmd.Output()
	if err != nil {
		log.Fatalf("\033[0;91m[ WHICH ]\033[0m :: Application %s not found", application)
	}
}
