package main

import (
	"log"
	"os"
	"text/template"
)

// tpl is a container holding all the parsed templates
var tpl *template.Template

func init() {
	// Must() does error checking and returns the template
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Name string
	Age  int
}

func main() {

	person := person{
		Name: "Gopher",
		Age:  4,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}
}
