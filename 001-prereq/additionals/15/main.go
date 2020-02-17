// Now add a method to type gator with this signature ...
//     greeting()
// … and have it print “Hello, I am a gator”.
// Create a value of type gator.
// Call the greeting() method from that value.

package main

import (
	"fmt"
)

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

var g1 gator = 42
var x int

func main() {
	fmt.Println("hey", g1)
	fmt.Printf("%T\n", g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	g1.greeting()

}
