// Now adding create user and delete user routes

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/imattf/gowebdev/042-mongodb/05-mongodb/02-update-all/models"
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
	//grab the id
	id := p.ByName("id")

	//Verify id is ObjectId hex representation, otherwise return not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//ObjectIdHex returns an ObjectId from the provided hex representation
	oid := bson.ObjectIdHex(id)

	//composite literal
	u := models.User{}

	//Fetch user
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
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

	//create the bson ID
	u.ID = bson.NewObjectId()

	//store the user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

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

	//grab the id
	id := p.ByName("id")

	//Verify id is ObjectId hex representation, otherwise return not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	//ObjectIdHex returns an ObjectId from the provided hex representation
	oid := bson.ObjectIdHex(id)

	//Delete the user
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //200
	fmt.Fprintf(w, "Deleted user: %s %s\n", oid, p.ByName("id"))

}
