// this is the same as my main04-betterDefaultServeMux.go

package main

import (
	"io"
	"net/http"
)

func hd(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark, bark")
}

func hc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow, meow, meow")
}

func main() {

	// type conversion to a Handler
	http.HandleFunc("/dog/", hd)
	http.HandleFunc("/cat", hc)

	http.ListenAndServe(":8080", nil)

}
