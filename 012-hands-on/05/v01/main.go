// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
// Using the data structure created, modify it to hold menu information for an unlimited number of restaurants.

package main

import (
	"log"
	"os"
	"text/template"
)

type menuItem struct {
	Number, Description string
	Price               float64
}

type menu struct {
	MenuName, MenuHours string
	MenuItems           []menuItem
}

type menus struct {
	Breakfast, Lunch, Dinner menu
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	m := menus{
		Breakfast: menu{
			MenuName:  "Breakfast",
			MenuHours: "8am-12pm",
			MenuItems: []menuItem{
				menuItem{"01", "eggs", 4.00},
				menuItem{"02", "toast", 5.00},
				menuItem{"03", "bacon", 6.00},
			},
		},
		Lunch: menu{
			MenuName:  "Lunch",
			MenuHours: "12pm-4pm",
			MenuItems: []menuItem{
				menuItem{"01", "fries", 4.25},
				menuItem{"02", "salad", 5.25},
				menuItem{"03", "sandwich", 6.25},
			},
		},
		Dinner: menu{
			MenuName:  "Dinner",
			MenuHours: "5pm-9pm",
			MenuItems: []menuItem{
				menuItem{"01", "salad", 4.75},
				menuItem{"02", "soup", 5.75},
				menuItem{"03", "steak", 6.75},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
