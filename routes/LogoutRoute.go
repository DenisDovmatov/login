package routes

import (
	"fmt"
	"github.com/martini-contrib/render"
	"loginB/session"
)

// Выход из пользователя
func LogoutHandler(rnd render.Render, s *session.Session) {
	if !s.Authorized{
		fmt.Print("You loginOFF you don't need logout")
		rnd.Redirect("/")
	}
	s.Authorized = false

	rnd.Redirect("/")
}