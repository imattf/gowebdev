// hand build session

package main

import (
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

//user Record
type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

// user storage
var dbUsers = map[string]user{} //userID, user

// session storage
var dbSessions = map[string]session{} //sessionID, userID
var dbSessionsCleaned time.Time

const sessionLength int = 30

// template prep
var tmpl *template.Template

func init() {

	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	//session mgt prep
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get the user
	u := getUser(w, req)

	//show session info
	showSessions()

	//show the index.html
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get the user
	u := getUser(w, req)

	// if user doesn't exist, redirect to /
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//must have right role to see /bar
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}

	//show sessions info
	showSessions()

	//show the bar.html
	tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {

	//see if already logged in
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		r := req.FormValue("role")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sid, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		//set the cookie on the response
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)

		//store the session info from the cookie in dbSessions
		dbSessions[cookie.Value] = session{un, time.Now()}

		//encrypt the password field
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		//store the user in dbUsers
		u = user{un, bs, fn, ln, r}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show session info
	showSessions()

	//show the signup.html
	tmpl.ExecuteTemplate(w, "signup.gohtml", u)

}

func login(w http.ResponseWriter, req *http.Request) {

	//see if already logged in
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		//is there a matching user name?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
			return
		}

		//create session
		sid, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		//set the cookie on the response
		cookie.MaxAge = sessionLength
		http.SetCookie(w, cookie)

		//store the session info from the cookie in dbSessions
		dbSessions[cookie.Value] = session{un, time.Now()}

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show session info
	showSessions()

	//show the login.html
	tmpl.ExecuteTemplate(w, "login.gohtml", u)

}

func logout(w http.ResponseWriter, req *http.Request) {
	// if not logged in, redirect to /
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//find the session cookie
	cookie, _ := req.Cookie("session")

	//delete the session from dbSession
	delete(dbSessions, cookie.Value)

	//remove the session cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	//set the cookie on the response
	http.SetCookie(w, cookie)

	//clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	//redirect
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
