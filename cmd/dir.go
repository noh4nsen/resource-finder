package cmd

import (
	"log"
	"os"
)

func CurrentContext() (path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return path
}

func Chdir(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Fatalf("\033[0;91m[ CHDIR ]\033[0m :: Path %s not found : %s", path, CurrentContext())
	}
}

func Mkdir(path string) {
	exists := os.Mkdir(path, 0755)
	if exists != nil {
		os.Chdir(path)
	}
	os.Chdir(path)
}
