// this is the same as my main04-betterDefaultServeMux.go
// ...just passing in a json value

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
	mux.GET("/user/:id", getUser) //new route + parm
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
