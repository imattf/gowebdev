// Building upon the code from the previous exercise:

// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

// Add code to respond to the following METHODS & ROUTES:
// GET /
// GET /apply
// POST /apply

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
			continue
		}

		// now handle multiple connections
		go serve(conn)

		fmt.Println("Code got here.")
	}

}

func serve(conn net.Conn) {
	// good housekeeping
	defer conn.Close()

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

		switch {
		case meth == "GET" && uri == "/":
			handleIndex(conn)
		case meth == "GET" && uri == "/apply":
			handleApply(conn)
		case meth == "POST" && uri == "/apply":
			handleApplyPost(conn)
		default:
			handleDefault(conn)
		}
	}
}

func handleIndex(conn net.Conn) {

	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
		`

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

}

func handleApply(conn net.Conn) {

	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET APPLY</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death? where's the poem?">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
		`

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

}

func handleApplyPost(conn net.Conn) {

	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
		`

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

}

func handleDefault(conn net.Conn) {

	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>"default"</h1>
		</body>
		</html>
		`

	// the status line
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

	// the reponse header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	// the message body
	io.WriteString(conn, body)

}
