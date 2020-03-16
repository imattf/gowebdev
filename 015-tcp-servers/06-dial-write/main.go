// use Dial to write on the connection to TCP Server
// run w/ /02-read/main.go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// bstring, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(bstring))

	fmt.Fprintln(conn, "I'm dialing in")
}
