// Initialize a MAP using a composite literal where the key is a string
// and the value is an int;
// print out the map;
// range over the map printing out just the key;
// range over the map printing out both the key and the value

package main

import "fmt"

func main() {
	myMap := map[string]int{
		"ann": 42,
		"bob": 41,
	}

	fmt.Println(myMap)

	for k, _ := range myMap {
		fmt.Println(k)
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
