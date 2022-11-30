package errpkg

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)


var Auth *auth
var Test *test
var General *general

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
	
    Auth = root.Auth 
    Test = root.Test 
    General = root.General 
}
