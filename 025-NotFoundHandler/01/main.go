// using http.NotFoundHandler for missing nice favico excepttions

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)

	//this appears to only make a difference in behavior w/ Chrome based browsers
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	fmt.Println("made it here")
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look in terminal")

}
