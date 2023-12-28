package cmd

import (
	"fmt"
	"log"
	"os"
	"resource-finder/model"
)

func WriteFile(script model.Script) {
	fileName := fmt.Sprintf("%s.%s.sh", script.Name, script.Directory)
	fileBody := []byte(script.Body)
	err := os.WriteFile(fileName, fileBody, 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
}
