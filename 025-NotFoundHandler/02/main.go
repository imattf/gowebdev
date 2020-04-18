// using http.NotFoundHandler for missing nice favico excepttions

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)

	//this appears to only make a difference in behavior w/ Chrome based browsers
	//http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {

		// if you comment out the call to http.NotFoundHandler above
		// this will get executed instead on Chrome based browser

		http.NotFound(w, req)
		return
	}

	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look in terminal")

}
