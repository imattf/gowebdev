// hand build session

package main

import (
	"fmt"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

//userRecord
type userRecord struct {
	UserName string
	First    string
	Last     string
}

// user storage
var dbUsers = map[string]userRecord{} //userID, user

// session storage
var dbSessions = map[string]string{} //sessionID, userID

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, create it
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
			// Secure: true,
			// HttpOnly: true,
			// Path:     "/",
		}
		//set the cookie on the response
		http.SetCookie(w, cookie)
	}
	fmt.Println("index method:", cookie)

	//if user exists already, get user
	var u userRecord
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	//process form submission; store the user & session info
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = userRecord{un, fn, ln}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	//show the index.html
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, redirect to???
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("bar method:", cookie)

	// retrieve userID info from session cookie
	userName, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	user := dbUsers[userName]
	//show the bar.html
	tmpl.ExecuteTemplate(w, "bar.gohtml", user)
}
