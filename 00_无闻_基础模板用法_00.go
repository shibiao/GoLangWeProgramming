package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

type Package struct{
	Name string
	NumFuncs int
	NumVars int
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl,err := template.New("go-web").Parse(`
Package name:{{.Name}}
Number of functions: {{.NumFuncs}}
Number of variables: {{.NumVars}}
`)
		if err != nil {
			fmt.Fprintf(writer,"Parse:%v",err)
			return
		}
		err = tmpl.Execute(writer,&Package{
			Name:"shibiao",
			NumFuncs:49,
			NumVars:90,
		})
		if err != nil {
			fmt.Fprintf(writer,"Exexute:%v",err)
			return
		}

	})
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Starting server...")
}