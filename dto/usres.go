package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	Success = "success"
	Error   = "error"
)

// Users slice to simulate a database
var Users []User

// User struct to define a user
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	UserId   int                `json:"userId,omitempty" bson:"userId,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Salt     string             `json:"salt,omitempty"`
	Name     string             `json:"name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Phone    string             `json:"phone,omitempty"`
	Location string             `json:"location,omitempty"`
}

// Give json for User struct to use in postman

// {
// 	"userId": 1,
// 	"name": "John Doe",
// 	"email": "
//
//
