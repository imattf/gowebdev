// Using cookies, track how many times a user has been to your website domain.

package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	// request the cookie file "my-cookie"
	// if cookie file doesn't exist (ErrNoCookie) then create it
	cookie, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	// extract current cookie Value field locally and increment
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	// update value in cookie
	http.SetCookie(w, cookie)
	// write cookie value to page
	io.WriteString(w, cookie.Value)

}
