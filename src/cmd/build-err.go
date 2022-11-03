package cmd

import (
	"github.com/spf13/cobra"
	"go/format"
	"gopkg.in/yaml.v3"
	"html/template"
	"io/ioutil"
	"log"
	"myclass_service/src/app"
	"myclass_service/src/config"
	"os"
	"strings"
)

type Field struct {
	Name   string
	Type   string
	YmlTag string
}

type WrapField struct {
	Name   string
	Fields []*Field
}

type WrapStruct struct {
	PackageName string
	Elements    []*WrapField
}

type PublicField struct {
	Name string
	Type string
}

type WrapPublicField struct {
	Elements []*PublicField
}

func buildErr(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "build-err",
		Run: func(cmd *cobra.Command, args []string) {
			structs, load := build()

			genStruct(structs)
			genLoad(load)
		},
	}
}

func build() (*WrapStruct, *WrapPublicField) {
	var tree map[string]map[string]app.Error

	f, err := os.Open("error.yml")

	if err != nil {
		log.Fatalln("fail open file", err)
	}

	defer func() {
		f.Close()
	}()

	err = yaml.NewDecoder(f).Decode(&tree)

	if err != nil {
		log.Fatalln("err parse tree data", err)
	}

	var wrapFields []*WrapField
	var rootErrorFields []*Field
	var publicFields []*PublicField

	for key, value := range tree {
		wrapFieldElement := WrapField{
			Name: key,
		}
		for name, _ := range value {
			field := Field{
				Name:   strings.Title(name),
				Type:   "*app.Error",
				YmlTag: name,
			}
			wrapFieldElement.Fields = append(wrapFieldElement.Fields, &field)
		}
		wrapFields = append(wrapFields, &wrapFieldElement)

		rootErrorField := Field{
			Name:   strings.Title(key),
			Type:   "*" + key,
			YmlTag: key,
		}
		rootErrorFields = append(rootErrorFields, &rootErrorField)

		publicField := PublicField{
			Name: strings.Title(key),
			Type: "*" + key,
		}
		publicFields = append(publicFields, &publicField)
	}

	wrapFields = append(wrapFields, &WrapField{
		Name:   "rootErr",
		Fields: rootErrorFields,
	})
	result := WrapStruct{
		PackageName: "err",
		Elements:    wrapFields,
	}

	return &result, &WrapPublicField{
		Elements: publicFields,
	}
}

func genStruct(structs *WrapStruct) {
	structFile := "src/packages/err/struct.go"

	f, err := os.Create(structFile)

	if err != nil {
		panic("err create file")
	}

	templ := template.Must(template.ParseFiles("src/templates/error/struct.tmpl"))
	templ.Execute(f, structs)
	f.Close()

	goFmt(structFile)

}

func goFmt(path string) {
	read, _ := ioutil.ReadFile(path)
	content, _ := format.Source(read)
	ioutil.WriteFile(path, []byte(content), 0)
}

func genLoad(loads *WrapPublicField) {
	loadFile := "src/packages/err/load.go"

	f, err := os.Create(loadFile)

	if err != nil {
		panic("err create file")
	}

	templ := template.Must(template.ParseFiles("src/templates/error/load.tmpl"))
	templ.Execute(f, loads)
	f.Close()

	//goFmt(loadFile)
}
