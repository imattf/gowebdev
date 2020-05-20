//Marshal and encode stuff..package main
//use Go data structure values for printing
// ...this has correct scope w/ Exported field names

package main

import (
	"encoding/json"
	"log"
	"os"
)

// will work... field names have to be Capitalized (Exported) for JSON to see it
type model struct {
	State    bool
	Pictures []string
}

// when you run this, should create a JSON
func main() {

	m := model{
		State: true,
		Pictures: []string{
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
