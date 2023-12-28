package main

import (
	"resource-finder/cmd"
	"resource-finder/core"
	"resource-finder/helper"
)

func main() {
	config := helper.LoadConfig("config.yaml")
	context := cmd.CurrentContext()
	core.WriteScripts(core.ProcessAllScripts(core.LoadReference(config), core.LoadDomains(config)), context)
}
