package db


type User struct {
	Id              string `bson:"_id"`
	FirstName           string
	LastName     string
	Email string
	Age string
	Password []byte
}
