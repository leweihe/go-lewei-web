package main

import (
	"log"
	"html/template"
	"os"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	log.Println("Start server...")
	template, err := template.New("go-lewei-web").Parse(
		"Package Name: {{.Name}} " +
			"Number of functions {{.NumFuncs}} " +
			"Number of variables {{.NumVars}} ")
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}

	err = template.Execute(os.Stdout, &Package{
		Name:     "go-lewei-web",
		NumFuncs: 12,
		NumVars:  1222,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
