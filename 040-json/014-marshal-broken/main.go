//Marshal and encode stuff..package main
//use Go data structure values for printing
// but has wrong scope

package main

import (
	"encoding/json"
	"log"
	"os"
)

// won't work... field names have to be Capitalized for JSON to see it
type model struct {
	state    bool
	pictures []string
}

// when you run this, nothing will come out
func main() {

	m := model{
		state: true,
		pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	//fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("error marshalling", err)
	}

	//fmt.Println(string(bs))
	os.Stdout.Write(bs)
}
