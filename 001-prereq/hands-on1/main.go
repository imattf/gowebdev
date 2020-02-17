//create a type square
//create a type circle
//attach a method to each that calculates area and returns it
//create a type shape which defines an interface as anything which has the area method
//create a func info which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle

package main

import (
	"fmt"
)

type square struct {
	side float64
}

type circle struct {
	rad float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.rad * c.rad
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{20}
	circ := circle{4}
	fmt.Println("Hello, playground", sq.area(), circ.area())
	info(sq)
	info(circ)
}
