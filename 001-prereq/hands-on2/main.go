// - create a struct that holds person fields
// - create a struct that holds secret agent fields and embeds person type
// - attach a method to person: pSpeak
// - attach a method to secret agent: saSpeak
// - create a variable of type person
// - create a variable of type secret agent
// - print a field from person
// - run pSpeak attached to the variable of type person
// - print a field from secret agent
// - run saSpeak attached to the variable of type secret agent
// - run pSpeak attached to the variable of type secret agent

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

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, `says, "I am a person"`)
}

func (su surfer) saSpeak() {
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

	fmt.Println(p1.fname)
	p1.pSpeak()

	fmt.Println(sa1.fname)
	sa1.saSpeak()
	sa1.person.pSpeak()

}
