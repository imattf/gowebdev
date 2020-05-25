// Now adding create user and delete user routes

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/imattf/gowebdev/042-mongodb/02-json/models"
	"github.com/julienschmidt/httprouter"
)

func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/user/:id", getUser)
	mux.POST("/user", createUser)       //new route
	mux.DELETE("/user/:id", deleteUser) //new route + parm
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/42">GO TO: http://localhost/user/42</a>
	</body>
	</html>
	`
	w.Header().Set("Context-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bondo",
		Gender: "mail",
		Age:    59,
		ID:     p.ByName("id"),
	}

	//Marshal into json
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//Write content-type, status-code, payload
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)

}

func createUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	//Decode the json that came in from the stream
	// ... put the json payload into u
	json.NewDecoder(req.Body).Decode(&u)

	//Change the ID
	u.ID = "007"

	//Now re-marshal back into a JSON and assign to a variable uj
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//Write content-type, status-code, payload
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)        // see the updated uj value

}

func deleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //200
	fmt.Fprintf(w, "Write code to delete user: %s\n", p.ByName("id"))

}
