package core

import (
	"log"
	"resource-finder/cmd"
	"resource-finder/model"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func LoadReference(config model.Config) (referenceStateList []model.ReferenceState) {
	reference, diags := tfconfig.LoadModule(config.Reference.Path)
	if diags != nil {
		log.Fatal(diags)
	}
	for _, workspace := range cmd.WorkspaceList(reference.Path) {
		var referenceState model.ReferenceState
		referenceState.LoadState(workspace, cmd.StateList(reference.Path, workspace))
		log.Printf("\033[0;94m[ %s ]\033[0m :: Loaded state references", workspace)
		referenceStateList = append(referenceStateList, referenceState)
	}
	return referenceStateList
}
