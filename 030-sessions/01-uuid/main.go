//using a different package from todd

package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, create it
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		//set the cookie on the response
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)

}
