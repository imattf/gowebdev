package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/imattf/gowebdev/042-mongodb/08-hands-on/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"

	"net/http"
)

//UserController is a...
type UserController struct {
	session map[string]models.User //use a map to users
}

//NewUserController is a...
func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

//GetUser is a...
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateUser is a...
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	id, _ := uuid.NewV4()
	u.Id = id.String()

	// store the user
	uc.session[u.Id] = u
	StoreUsers(uc.session) //update the json file

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser is a...
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	//Delete a user from a map, based on a key value
	delete(uc.session, id)
	StoreUsers(uc.session) //update the json file

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", id, "\n")
}

//StoreUsers is a...
func StoreUsers(m map[string]models.User) {

	//create the json file
	f, err := os.Create("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//write to the json file
	json.NewEncoder(f).Encode(m)
}

//LoadUsers is a...
func LoadUsers() map[string]models.User {

	//make the local map store of users
	m := make(map[string]models.User)

	//open the json file
	f, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	//write to the json file
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}
