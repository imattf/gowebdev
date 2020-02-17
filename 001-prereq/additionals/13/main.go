//Adding onto this code: Can you assign “g1” to “x”? Why or why not?

package main

import (
	"fmt"
)

type gator int

var g1 gator = 42
var x int

func main() {
	fmt.Println("hey", g1)
	fmt.Printf("%T\n", g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = g1
	fmt.Println(x)
	// you CANNOT assign g1 to x. They are two different types!

}
