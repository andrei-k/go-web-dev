package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// tpl is a container holding all the parsed templates
var tpl *template.Template

// Create a FuncMap to register functions
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a custom func
// "ft" slices a string and returns the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	// Must() does error checking and returns the template
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
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
