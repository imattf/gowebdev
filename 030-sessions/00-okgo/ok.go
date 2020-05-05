package main

import "fmt"

func main() {
	//this m := map[string]int{}, will do this...
	m := make(map[string]int)
	//m["me"] = 42
	if meAge, ok := m["me"]; !ok {
		fmt.Println(meAge, ok)
	}
}
