//unmarshal and encode stuff..package main
//use Go data structure values for printing

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data []string
	recvd := `null`

	err := json.Unmarshal([]byte(recvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	fmt.Println(data)
	fmt.Println(len(data))
	fmt.Println(cap(data))
}
