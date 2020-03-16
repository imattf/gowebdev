// TCP server - write to connection
// We can create our own tcp server using the net package from the standard library.
// There are three main steps: (1) Listen (2) Accept (3) Write or Read to the connection.
// We will use telnet to call into the TCP server we created.
// Telnet provides bidirectional interactive text-oriented communication
// using a virtual terminal connection over the Transmission Control Protocol (TCP).

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	// create the server...
	svr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// close the server
	defer svr.Close()

	// start a telnet connection pipe
	for {
		conn, err := svr.Accept()
		if err != nil {
			log.Println(err)
		}

		//execute on the telnet session
		io.WriteString(conn, "Hello ")
		io.WriteString(conn, "World\n")
		fmt.Fprintln(conn, "this is Frptintln")
		fmt.Fprintf(conn, "%v", "This is Fprintf\n")

		//does fmt.Println work? ~ it writes to stdout
		fmt.Println("nope")

		//close the telnet session
		conn.Close()

	}
}
