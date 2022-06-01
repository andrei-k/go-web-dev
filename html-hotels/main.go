package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := hotels{
		hotel{
			Name:    "AAA",
			Address: "111",
			City:    "Los Angeles",
			Zip:     "90210",
			Region:  "Southern",
		},
		hotel{
			Name:    "BBB",
			Address: "222",
			City:    "San Francisco",
			Zip:     "91890",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
