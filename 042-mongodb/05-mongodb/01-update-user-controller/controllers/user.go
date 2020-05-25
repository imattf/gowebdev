// Now adding create user and delete user routes

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/imattf/gowebdev/042-mongodb/05-mongodb/01-update-user-controller/models"
	"github.com/julienschmidt/httprouter"
)

//UserController is a method to ...
type UserController struct {
	session *mgo.Session
}

//NewUserController is a function to ...
// ...convenience function to create a new instance
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//GetUser is a method to ...
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
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

//CreateUser is a method to ...
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
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

//DeleteUser is a method to ...
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //200
	fmt.Fprintf(w, "Write code to delete user: %s\n", p.ByName("id"))

}
