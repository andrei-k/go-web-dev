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

func main() {
	names := map[string]string{
		"1": "Mary",
		"2": "Bob",
		"3": "Sam",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
