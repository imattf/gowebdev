//example of functions to handle http (i.e. request & response) in tcp server
//create a multiplexer (mux, servemux, router, server, http mux, ...)
// so that your server responds to different URIs and METHODS.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	// read request
	request(conn)

	// write response
	//respond(conn)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// if request line, send connection and request line to multiplexer
			multiplexer(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func multiplexer(conn net.Conn, ln string) {
	// request line
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("***METHOD", method)
	fmt.Println("***URI", uri)
	//try camel case too; Case matters
	if method == "GET" && uri == "/dog" {
		dog(conn)
	}
	if method == "GET" && uri == "/cat" {
		cat(conn)
	}
	if method == "GET" && uri == "/" {
		index(conn)
	}
}

// func respond(conn net.Conn) {
// 	//all this stuff getting sent back to the client as a file...

// 	body :=
// 		`<!DOCTYPE html>
// 		<html lang="en"><head><meta charset="UTF-8"><title></title></head>
// 			<body>
// 			<strong>Hello, World</strong>
// 			</body>
// 		</html>`

// 	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
// 	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
// 	fmt.Fprint(conn, "Content-Type: text/html\r\n")
// 	fmt.Fprint(conn, "\r\n")
// 	fmt.Fprint(conn, body)

// }

func dog(conn net.Conn) {
	body :=
		`<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title></title></head>
		<body>
		<a href="/">index</a><br>
		<strong>Dog...</strong><br>
		Bark, Bark!!
		</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func cat(conn net.Conn) {
	body :=
		`<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title></title></head>
		<body>
		<a href="/">index</a><br>
		<strong>Cat...</strong><br>
		Meow, Meow!!
		</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func index(conn net.Conn) {
	body :=
		`<!DOCTYPE html>
		<html lang="en"><head><meta charset="UTF-8"><title></title></head>
			<body>
			<strong>Index...</strong>
			<a href="/">index</a><br>
			<a href="/dog">Dog</a><br>
			<a href="/cat">Cat</a><br>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
