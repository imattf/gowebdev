package session

import (
	"github.com/imattf/gowebdev/042-mongodb/10-hands-on/models"

	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Length is a...
const Length int = 30

//Users is a...
var Users = map[string]models.User{} // user ID, user
//Sessions is a...
var Sessions = map[string]models.Session{} // session ID, session
//LastCleaned is a...
var LastCleaned time.Time = time.Now()

//GetUser is a ...
func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	ck, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	ck.MaxAge = Length
	http.SetCookie(w, ck)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
		u = Users[s.UserName]
	}
	return u
}

//AlreadyLoggedIn is a ...
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}
	_, ok = Users[s.UserName]
	// refresh session
	ck.MaxAge = Length
	http.SetCookie(w, ck)
	return ok
}

//CleanSessions is a ...
func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// ShowSessions is a for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
