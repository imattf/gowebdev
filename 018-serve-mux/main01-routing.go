package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("do what you want to do")
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "bark, bark, bark")
	case "/cat":
		io.WriteString(res, "meow, meow, meow")
	}
}

func main() {
	var d hotdog
	// can you print the value of d to the client later?
	d = 42
	http.ListenAndServe(":8080", d)

}
