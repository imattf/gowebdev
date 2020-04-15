// Building upon the code from the previous exercise:

// Add code to WRITE to the connection.

package main

import (
	"bufio"
	"fmt"
	"io"
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
		// scanner := bufio.NewScanner(conn)
		// for scanner.Scan() {
		// 	ln := scanner.Text()
		// 	fmt.Println(ln)
		// 	if ln == "" {
		// 		//when ln is empty, the header is done
		// 		fmt.Println("THUS, END OF HTTP REQUEST HEADERS")
		//      break
		// 	}
		// }

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

	io.WriteString(conn, "I see you connected")

	defer conn.Close()

}
