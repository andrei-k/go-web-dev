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
				{"CSCI-40", "Introduction to Programming in Go", "4"},
				{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming with Go", "5"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, year)
	if err != nil {
		log.Fatalln(err)
	}
}
