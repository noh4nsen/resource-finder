package core

import (
	"resource-finder/cmd"
	"resource-finder/model"
)

func WriteScripts(moveScripts []model.Script, context string) {
	for _, script := range moveScripts {
		cmd.Chdir(context)
		cmd.Mkdir(script.Directory)
		cmd.WriteFile(script)
	}
}
