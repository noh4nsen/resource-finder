package model

import "fmt"

type Script struct {
	Directory string
	Name      string
	Body      string
}

func (moveScript *Script) Initialize(directory, workspace string) {
	moveScript.Directory = directory
	moveScript.Name = workspace
	moveScript.loadScriptBodyHeader(workspace)
}

func (moveScript *Script) loadScriptBodyHeader(workspace string) {
	moveScript.Body = moveScript.Body + "#!/bin/bash\n\n"
	moveScript.Body = moveScript.Body + "terraform workspace select " + workspace + "\n\n"
}

func (moveScript *Script) LoadScriptBody(formattedTime, directory, workspace, referenceResource, resourceSuffix, moduleName, moduleResource string) {
	if resourceSuffix == "GENERAL" {
		moveScript.Body = moveScript.Body + fmt.Sprintf("terraform state mv --state-out '%s.%s.%s.migrated.tfstate' %s %s%s\n",
			formattedTime, moveScript.Directory, moveScript.Name,
			referenceResource, moduleName, moduleResource)
	} else {
		moveScript.Body = moveScript.Body + fmt.Sprintf("terraform state mv --state-out '%s.%s.%s.migrated.tfstate' %s[\\\"%s\\\"] %s%s[\\\"%s\\\"]\n",
			formattedTime, moveScript.Directory, moveScript.Name,
			referenceResource, resourceSuffix,
			moduleName, moduleResource, resourceSuffix)
	}
}
