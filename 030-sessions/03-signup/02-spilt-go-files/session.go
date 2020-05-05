package main

import "net/http"

func getUser(req *http.Request) userRecord {

	var u userRecord

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, redirect to???
	if err != nil {
		return u
	}

	//if user exists already, get user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u

}

func alreadyLoggedIn(req *http.Request) bool {

	//get cookie
	cookie, err := req.Cookie("session")

	// if it doesn't exist, return "not logged in"
	if err != nil {
		return false
	}

	//if user exists already, get user
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok

}
