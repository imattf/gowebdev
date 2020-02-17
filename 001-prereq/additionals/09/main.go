// Give a method to both the “truck” and “sedan” types with the following signature
//     transportationDevice() string
// Have each func return a string saying what they do.
// Create a value of type truck and populate the fields.
// Create a value of type sedan and populate the fields.
// Call the method for each value.

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

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I haul things")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I haul people")
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
	fmt.Println(myTruck.transportationDevice())

	fmt.Println(myCar)
	fmt.Println(myCar.luxury)
	fmt.Println(myCar.transportationDevice())

}
