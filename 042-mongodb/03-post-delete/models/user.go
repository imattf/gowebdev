package models

//User is the model for this app
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	ID     string `json:"id"`
}
