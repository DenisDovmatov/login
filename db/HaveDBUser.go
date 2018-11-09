package db

import (
	"fmt"
)
var i int64

func HaveDBUser(email string)(bool, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email=$1", email)
	fmt.Println(row)
	user := new(User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Password)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Этот почтовый адрес свободен")
		return false, err
	}
	fmt.Println("Этот адрес уже используется")
	return true, nil
}
