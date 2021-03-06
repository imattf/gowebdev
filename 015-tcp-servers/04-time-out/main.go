//example of telnet connection read from & write to client
// TCP server - read from & write to connection
// Now we are going to read and write from/to our connection.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	s := time.Now().Add(10 * time.Second)
	fmt.Println(s)
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// we now get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here, really")
}
