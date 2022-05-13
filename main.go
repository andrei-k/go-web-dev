package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// tpl is a container holding all the parsed templates
var tpl *template.Template

func init() {
	// Must() does error checking and returns the template
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func monthDayYear(t time.Time) string {
	// Time.Format accepts a layout in this syntax: 01/02 03:04:05PM '06 -0700
	// In my case, I want:
	// day (referred to as 02)
	// month (referred to as 01)
	// year (referred to as 06)
	return t.Format("02-01-2006")
}

func convertString(s string) string {
	return s + "!"
}

// Create a FuncMap to register functions
var fm = template.FuncMap{
	"fdateMDY":      monthDayYear,
	"convertString": convertString,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
