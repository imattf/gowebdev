// Building upon the code from the previous exercise:

// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function
//  called "serve".
// Pass the connection of type net.Conn as an argument into this function.
// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go serve(conn)

		fmt.Println("Code got here.")
	}

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			//when ln is empty, the header is done
			fmt.Println("THUS, END OF HTTP REQUEST HEADERS")
			break
		}
	}

	defer conn.Close()

}
