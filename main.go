package main

import (
	"html/template"
	"loginB/db"

	"loginB/routes"
	"loginB/session"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	infoDB = "user=beli password=w111 dbname=logindb sslmode=disable"
)


func unescape(x string) interface{} {
	return template.HTML(x)
}


func main() {

	db.InitDB(infoDB)

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Use(session.Middleware)

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "main",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))

	staticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", staticOptions))

	m.Get("/", routes.IndexHandler)
	m.Get("/profile", routes.ProfileHandler)
	m.Get("/register", routes.GetRegisterHandler)
	m.Post("/register", routes.PostRegisterHandler)
	m.Get("/login", routes.GetLoginHandler)
	m.Post("/login", routes.PostLoginHandler)
	m.Get("/logout", routes.LogoutHandler)
	m.Post("/have", routes.HaveEmailHandler)
	m.RunOnAddr("127.0.0.53:3004")
}
