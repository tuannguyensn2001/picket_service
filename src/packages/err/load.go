package errpkg

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)


var General *general
var Auth *auth

func LoadError() {
	root := rootErr{}

	file, err := os.ReadFile("error.yml")

	if err != nil {
		log.Fatalln("error load error", err)
	}

	err = yaml.Unmarshal(file, &root)

	if err != nil {
		log.Fatalln("error unmarshal file", err)
	}

	//General = root.General
	
    General = root.General 
    Auth = root.Auth 
}
