package models

type User struct {
	Id              string
	FirstName           string
	LastName     string
	Email string
	Age string
	Password []byte
}

func NewUser(id, firstname, lastname, email, age string, password []byte) *User {
	return &User{id, firstname, lastname, email, age, password}
}
