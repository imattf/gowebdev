// Initialize a SLICE of int using a composite literal; print out the slice;
// range over the slice printing out just the index;
// range over the slice printing out both the index and the value

package main

import "fmt"

func main() {
	mySlice := []int{41, 42, 43}
	for i, _ := range mySlice {
		fmt.Println(i)
	}
	for i, x := range mySlice {
		fmt.Println(i, x)
	}

}
