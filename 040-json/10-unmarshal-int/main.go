//unmarshal and encode stuff..package main
//use Go data structure values for printing

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data int
	recvd := `42`

	err := json.Unmarshal([]byte(recvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	fmt.Println(data)
}
