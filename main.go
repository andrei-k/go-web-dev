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

	p1 := person{
		Name: "Gopher",
		Age:  4,
	}

	p2 := person{
		Name: "Mary",
		Age:  40,
	}

	p3 := person{
		Name: "John",
		Age:  18,
	}

	people := []person{p1, p2, p3}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}
