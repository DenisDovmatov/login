package routes

import (
	"fmt"
	"loginB/db"
	"net/http"
	"github.com/martini-contrib/render"
	"loginB/session"
)

// Регистрация пользователя
func GetRegisterHandler(rnd render.Render, s *session.Session) {
	if s.Authorized{
		fmt.Print("You loginON")
		rnd.Redirect("/profile")
	}
	rnd.HTML(200, "register", nil)
}

// Отправка данных регистрации на сервер
func PostRegisterHandler(rnd render.Render, r *http.Request, s *session.Session) {
	if !s.Authorized{
		fmt.Print("Вы в гостевом режиме сейчас начнется регистрация")
	}else{
		fmt.Println("Вы уже залогинены выйдете из профиля чтобы зарегестрировать новый аккаунт")
		rnd.Redirect("/profile")
	}

	first := r.FormValue("first")
	last := r.FormValue("last")
	email := r.FormValue("email")
	age := r.FormValue("age")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")

	// Проверка значений

	if first == "" || last == "" || email == "" || age == "" || password == "" || password2 == ""{
		fmt.Println("Введенно не корректное значение")
		rnd.Error(400)
		return
	}
	// Сравнение паролей
	if password != password2{
		fmt.Println("Пароли не совпадают")
		rnd.Error(400)
	}
	// Проверка почты
	have, err := db.HaveDBUser(email)

	if have{

		message := "Эта почта уже зарегестрирована"
		fmt.Println(message)

		rnd.HTML(200, "register", s)
		return
	}

	// Сохранение нового пользователя в базу
	res, err := db.SaveDBUser(first, last, email, age, password)
	if err != nil {
		rnd.Error(500)
		return
	}else{
		// Вывод в консоль при создании пользователя
		fmt.Printf("Пользователь %s %s создан (%d)\n", first, last, res)
		// Запись в куки
		s.First = first
		s.Last = last
		s.Email = email
		s.Age = age
		s.Authorized = true
	}

	rnd.Redirect("/")
}
