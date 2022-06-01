package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

// This method gets passed to the template as well
func (c course) Combine() string {
	return c.Number + " / " + c.Name + " / " + c.Units
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	year := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CS-50", "Course 1", "4"},
				{"CS-100", "Course 2", "4"},
				{"CS-150", "Course 3", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CS-200", "Course 4", "5"},
				{"CS-250", "Course 5", "5"},
				{"CS-300", "Course 6", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, year)
	if err != nil {
		log.Fatalln(err)
	}
}
