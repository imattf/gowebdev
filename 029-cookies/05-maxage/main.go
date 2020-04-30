// passing data
// delete a cookie w/ MaxAge

package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)

}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c1)

}

func expire(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	c1.MaxAge = -1 //delete the cooke
	http.SetCookie(w, c1)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
