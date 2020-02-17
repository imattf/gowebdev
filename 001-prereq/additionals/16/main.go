// Adding onto this code: create another type called “flamingo”.
// Make the underlying type of “flamingo” bool.
// Give “flamingo” a method with this signature …
//     greeting()
// … and have it print “Hello, I am pink and beautiful and wonderful.”
// Now create a new type “swampCreature”.
// The underlying type of “swapCreature” is interface.
// The behavior which the “swampCreature” interface defines is such that
// any type which has this method …
//     greeting()
// … will implicitly implement the “swampCreature” interface.
// Create a func called “bayou” which takes a value of type “swampCreature” as an argument.
// Have this func print out the greeting for whatever “swampCreature” might be passed in

package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	var f1 flamingo

	bayou(g1)
	bayou(f1)
}
