package helper

import (
	"log"
	"os"
	"resource-finder/model"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string) (config model.Config) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
