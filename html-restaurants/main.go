package main

import (
	"html/template"
	"log"
	"os"
)

type Item struct {
	Name  string
	Price float64
}

type Meal struct {
	Name  string
	Items []Item
}

type Menu []Meal

type Restaurant struct {
	Name string
	Menu Menu
}

type Restaurants []Restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := Restaurants{
		Restaurant{
			Name: "Tandem",
			Menu: []Meal{
				Meal{
					Name: "Breakfast",
					Items: []Item{
						{"Bacon", 3.99},
						{"Eggs", 1.99},
					},
				},
				Meal{
					Name: "Lunch",
					Items: []Item{
						{"Hamburger", 7.99},
						{"Salad", 8.99},
					},
				},
				Meal{
					Name: "Dinner",
					Items: []Item{
						Item{
							Name:  "Chicken",
							Price: 9.99,
						},
						Item{
							Name:  "Pork",
							Price: 7.99,
						},
					},
				},
			},
		},
		Restaurant{
			Name: "Kitchen Gallery",
			Menu: []Meal{
				Meal{
					Name: "Breakfast",
					Items: []Item{
						{"Sausage", 3.99},
						{"Toast", 1.99},
					},
				},
				Meal{
					Name: "Lunch",
					Items: []Item{
						{"Fish", 7.99},
						{"Soup", 8.99},
					},
				},
				Meal{
					Name: "Dinner",
					Items: []Item{
						Item{
							Name:  "Steak",
							Price: 9.99,
						},
						Item{
							Name:  "Chili",
							Price: 7.99,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
