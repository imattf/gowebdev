//review methods, interfaces and polymorphism

package main

import "fmt"

type person struct {
	fname string
	lname string
}

type surfer struct {
	person
	licenseToSurf bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "I am a person"`)
}

func (su surfer) speak() {
	fmt.Println(su.fname, su.lname, `says, "Surf's Up!!"`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Bob",
		"Jones",
	}

	sa1 := surfer{
		person{
			"Kon",
			"Tiki",
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)
}
