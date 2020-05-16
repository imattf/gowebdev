// photo-blog
// store some user data w/ cookie; session uuid which will serve as the session ID when we need it

package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get cookie
	cookie := getCookie(w, req)

	tmpl.ExecuteTemplate(w, "index.gohtml", cookie)
}

//getCookie function out here
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session")
	// if the session cookie doesn't exist, create it
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		//set the cookie on the response
		http.SetCookie(w, cookie)
	}
	return cookie
}
