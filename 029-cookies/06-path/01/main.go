//goofy1

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users/goofy", set)
	http.HandleFunc("/users/goofy/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// find cookie
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Println(c, err)
	fmt.Fprintln(w, "FINDING YOUR COOKIE:", c)

	//so this cant find cookie after its been created, hmmm
	//...because the path is /users on the set. See 02 for solution
}

func set(w http.ResponseWriter, req *http.Request) {

	// create the cookie
	c := &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	}
	// set the cookie on the browser; add cookie header to response writer's headers
	http.SetCookie(w, c)
	fmt.Println(c)
	fmt.Fprintln(w, "SETTING YOUR COOKIE: ", c)
	fmt.Fprintln(w, "...go to dev tools cookies", c)

}

func read(w http.ResponseWriter, req *http.Request) {

	// find cookie
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "READING YOUR COOKIE:", c)

}
