package db

import (
	"fmt"
)

func GetDBUser(email string)(*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email=$1", email)
	fmt.Println(row)

	user := new(User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
