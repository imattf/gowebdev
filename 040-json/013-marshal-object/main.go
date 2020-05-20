//unmarshal and encode stuff..package main
//use Go data structure values for printing

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {

	m := model{}

	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("error marshalling", err)
	}

	fmt.Println(string(bs))
}
