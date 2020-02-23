// Create a data structure to pass to a template which
// contains information about California hotels including
// Name, Address, City, Zip, Region
// region can be: Southern, Central, Northern
// can hold an unlimited number of hotels

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	// stuff goes here
	// hotel1 := hotel{
	// 	Name:    "Tioga",
	// 	Address: "123 Main St",
	// 	City:    "Merced CA",
	// 	Zip:     "789",
	// 	Region:  "Central",
	// }

	// hotel2 := hotel{
	// 	Name:    "Coronado",
	// 	Address: "123 Main St",
	// 	City:    "Coronado CA",
	// 	Zip:     "456",
	// 	Region:  "Southern",
	// }

	// hotel3 := hotel{
	// 	Name:    "Palace",
	// 	Address: "123 Main St",
	// 	City:    "San Francisco CA",
	// 	Zip:     "456",
	// 	Region:  "Northern",
	// }

	// hotel4 := hotel{
	// 	Name:    "Palace",
	// 	Address: "123 Main St",
	// 	City:    "San Francisco CA",
	// 	Zip:     "456",
	// 	Region:  "Northern",
	// }

	// caHotels := []hotel{hotel1, hotel2, hotel3, hotel4}

	h := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "Southern",
		},
		hotel{
			Name:    "Hotel Coronado",
			Address: "123 Main St",
			City:    "Coronaso",
			Zip:     "95612",
			Region:  "Southern",
		},
	}

	err := tmpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
