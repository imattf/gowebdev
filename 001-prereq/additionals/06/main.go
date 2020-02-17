// Add a method to type “person” with the identifier “walk”.
// Have the func return this string: <person’s first name> is walking.
// Remember, you make a func a method by giving the func a receiver.
//     func (r receiver) identifier(parameters) (returns) {
//         <code>
//     }
// To return a string, use fmt.Sprintln.
// Call the method assigning the returned string to a variable with the identifier “s”.
// Print out “s”.

package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	myPerson := person{
		fName:   "Ann",
		lName:   "Banana",
		favFood: []string{"apples", "peaches"},
	}
	s := myPerson.walk()
	fmt.Println(s)

}
