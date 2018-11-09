package db

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func SaveDBUser(first string, last string, email string, age string, password string) (int64, error) {

	fmt.Println( "Запись в базу пользователя----(", first, last, email, age, password, ")")

	id, _ := uuid.NewV4()

	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error("Сервер для криптования пароля не доступен")
	}

	result, err := db.Exec("INSERT INTO users VALUES($1, $2, $3, $4, $5, $6)",
																	id, first, last, email, age, passwordCrypt)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	row, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return row, nil
}
