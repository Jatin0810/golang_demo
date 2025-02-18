package model

type UserModel struct {
	Firstname    string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Age          int    `json:"age,omitempty" bson:"age,omitempty"`
	MobileNumber int `json:"mobilenumber,omitempty" bson:"mobilenumber,omitempty"`
}
