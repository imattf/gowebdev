// Building upon the code from the previous exercise:

// Before we WRITE our RESPONSE, let's WRITE to our RESPONSE the STATUS LINE and some RESPONSE HEADERS.
// Remember the request line and status line:
// 	REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF https://tools.ietf.org/html/rfc7230#section-3.1.1
// 	RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2
// Write the following strings to the response - use io.WriteString for all of the following except the second and third:
// 	"HTTP/1.1 200 OK\r\n"
// 	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
// 	fmt.Fprint(c, "Content-Type: text/plain\r\n")
// 	"\r\n"
// Look in your browser "developer tools" under the network tab.
// Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.

// and now...
// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
// Add this data to your REPONSE so that this data is displayed in the browser.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

		serve(conn)

		fmt.Println("Code got here.")
	}

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0

	var meth, uri string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line is line 0
			meth = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("***METHOD", meth)
			fmt.Println("***URI", uri)

			//todd's way...
			// xs := strings.Fields(ln)
			// meth = xs[0]
			// uri = xs[1]
			// fmt.Println("***METHOD", meth)
			// fmt.Println("***URI", uri)
		}

		if ln == "" {
			//when ln is empty, the request header is done
			fmt.Println("THUS, END OF HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	// the body construct w/ concatination
	body := "THIS IS THE PAYLOAD"
	body += "\n"
	body += meth
	body += "\n"
	body += uri

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

	defer conn.Close()

}
