// passing data
// using multiple cookie files

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in browser, go to developer tools, cookies")

	// write to console...
	// fmt.Print("Your request method at foo: ", req.Method, "\n\n")

}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("my-cookie")
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		// return
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		// return
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		// return
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3", c3)
	}

}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other general value",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other specific value",
	})

	fmt.Fprintln(w, "ABUNDANT COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in browser, go to developer tools, cookies")
}
