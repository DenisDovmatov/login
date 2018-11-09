package routes

import (
	"github.com/martini-contrib/render"
	"loginB/session"
)

// Начальная страница
func IndexHandler(rnd render.Render, s *session.Session) {
	model := s
	rnd.HTML(200, "index", model)
}