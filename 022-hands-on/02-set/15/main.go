// Building upon the code from the previous exercise:

// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD"
// (and everything else it contained: request method, request URI) to an HTML PAGE
// that prints "HOLY COW THIS IS LOW LEVEL" in <h1> tags.

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
	// body := "THIS IS THE PAYLOAD"
	// body += "\n"
	// body += meth
	// body += "\n"
	// body += uri

	body := `<!DOCTYPE html>
			<html lang="en">
			<head><meta charset="UTF-8">
			<title></title>
			</head>
			<body>
			<h1>HOLY COW THIS IS LOW LEVEL</h1>
			<h1>` + meth +
		`</h1>
			<h1>` + uri +
		`</h1>
			</body>
			</html>`

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

	defer conn.Close()

}
