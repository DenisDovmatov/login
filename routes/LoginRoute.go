package routes

import (
	"fmt"
	"loginB/db"
	"net/http"
	"github.com/martini-contrib/render"
	"loginB/session"
	bcrypt
)

//  Страница логина
func GetLoginHandler(rnd render.Render, s *session.Session) {
	// Проверка авторизации
	if s.Authorized{
		fmt.Print("You loginON")
		rnd.Redirect("/")
	}
	rnd.HTML(200, "login", s)

}

func PostLoginHandler(rnd render.Render, r *http.Request, s *session.Session, w http.ResponseWriter) {
	// Проверка авторизации
	if s.Authorized{
		fmt.Println("Вы уже залогинены")
		rnd.Redirect("/")
	}

	// Получение значений
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Проверка пользователя по базе
	user, err := db.GetDBUser(email)

	if err != nil {
		http.Error(w, "Такая почта не зарегестрирована", http.StatusForbidden)
		return
	}
	// Проверка пароля
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		http.Error(w, "Вы ввели не корректный пароль", http.StatusForbidden)
		return
	}
	// Запись куки
	s.First = user.FirstName
	s.Last = user.LastName
	s.Email = user.Email
	s.Age = user.Age
	s.Authorized = true

	rnd.Redirect("/")
}
