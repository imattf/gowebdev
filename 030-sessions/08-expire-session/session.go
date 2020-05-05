package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {

	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	//if user exists already, get user
	var u user
	if s, ok := dbSessions[cookie.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
		u = dbUsers[s.un]
	}
	return u

}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, return "not logged in"
	if err != nil {
		return false
	}

	//if user exists already, get user
	s, ok := dbSessions[cookie.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}
	_, ok = dbUsers[s.un]

	//refresh the session
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok

}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	fmt.Println("AFTER CLEAN")
	showSessions()

}

func showSessions() {
	fmt.Println("****************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("   ")
}
