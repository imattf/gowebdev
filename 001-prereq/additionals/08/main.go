// Using the vehicle, truck, and sedan structs:
// using a composite literal,
// create a value of type truck and assign values to the fields;
// using a composite literal,
// create a value of type sedan and assign values to the fields.
// Print out each of these values. Print out a single field from each of these values.

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	myTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	myCar := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(myTruck)
	fmt.Println(myTruck.vehicle.doors)

	fmt.Println(myCar)
	fmt.Println(myCar.luxury)

}
