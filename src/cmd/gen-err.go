package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
	"picket/src/app"
	"picket/src/config"
	"strconv"
	"strings"
)

type tmp struct {
	Fields []Parent
}

type Child struct {
	Name       string
	Message    string
	StatusCode int
	Code       int
	GrpcCode   int
}

type Parent struct {
	Name     string
	Children []Child
}

func genError(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "gen-error",
		Run: func(cmd *cobra.Command, args []string) {

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Nhap category: ")
			parent, err := reader.ReadString('\n')
			parent = strings.TrimSuffix(parent, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap title: ")
			child, err := reader.ReadString('\n')
			child = strings.TrimSuffix(child, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap message: ")
			message, err := reader.ReadString('\n')
			message = strings.TrimSuffix(message, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap status: ")
			statusCodeStr, err := reader.ReadString('\n')
			statusCodeStr = strings.TrimSuffix(statusCodeStr, "\n")
			statusCode, err := strconv.Atoi(statusCodeStr)
			if err != nil {
				fmt.Println("status code is not valid", err)
			}

			fmt.Print("Nhap grpc code: ")
			grpcCodeStr, err := reader.ReadString('\n')
			grpcCodeStr = strings.TrimSuffix(grpcCodeStr, "\n")
			grpcCode, err := strconv.Atoi(statusCodeStr)
			if err != nil {
				fmt.Println("grpc code is not valid", err)
			}

			m, err := app.LoadErr("./error.yml")
			if err != nil {
				fmt.Print("fail load err", err)
			}

			max := 0

			for _, val := range m {
				for _, val := range val {
					if max < val.Code {
						max = val.Code
					}
				}
			}

			nextCode := max + 1

			httpError := app.Error{
				StatusCode: statusCode,
				Message:    message,
				Code:       nextCode,
				GrpcCode:   grpcCode,
			}

			_, ok := m[parent]
			if !ok {
				m[parent] = make(map[string]app.Error)
			}

			m[parent][child] = httpError

			file := "error.yml"

			f, err := os.Create(file)
			if err != nil {
				fmt.Println("err create file", err)
			}

			parentFields := make([]Parent, 0)

			for k, v := range m {
				item := Parent{
					Name: k,
				}
				children := make([]Child, 0)
				for k1, v1 := range v {
					item2 := Child{
						Name:       k1,
						Message:    v1.Message,
						StatusCode: v1.StatusCode,
						Code:       v1.Code,
						GrpcCode:   v1.GrpcCode,
					}
					children = append(children, item2)
				}
				item.Children = children
				parentFields = append(parentFields, item)
			}

			templ := template.Must(template.ParseFiles("src/templates/error/error.tmpl"))
			templ.Execute(f, tmp{
				Fields: parentFields,
			})
			f.Close()

			//deleteLine(file)
		},
	}
}
