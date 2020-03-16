// TCP server - read from connection using bufio.Scanner
// We will now modify our TCP server to handle multiple connections.
// We will do this by using goroutines. We will also modify our TCP server to read from the connection.
// We will then contact our TCP server on port 8080 using our web browser.
// This will allow us to see the text sent from the browser to the TCP server and how this text adheres to HTTP (RFC 7230).

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
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")

	//...but if you send another http request the above code will be executed
}
