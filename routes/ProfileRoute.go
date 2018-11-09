package routes

import (
	"fmt"
	"github.com/martini-contrib/render"
	"loginB/session"
)

// Профиль юзера
func ProfileHandler(rnd render.Render, s *session.Session) {
	if !s.Authorized{
		fmt.Print("You loginOFF, Now you go in loginForm")
		rnd.Redirect("/login")
	}

	model := s

	rnd.HTML(200, "profile", model)
}