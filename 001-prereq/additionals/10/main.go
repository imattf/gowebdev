// Create a new type called “transportation”.
// The underlying type of this new type is interface.
// An interface defines functionality.
// Said another way, an interface defines behavior.
// For this interface, any other type that has a method with this signature …
//     transportationDevice() string
// … will automatically, implicitly implement this interface.
// Create a func called “report” that takes a value of type “transportation” as an argument.
// The func should print the string returned by “transportationDevice()” being called
// for any type that implements the “transportation” interface.
// Call “report” passing in a value of type truck.
// Call “report” passing in a value of type sedan

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

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I haul things")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I haul people")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
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

	report(myTruck)
	report(myCar)

}
