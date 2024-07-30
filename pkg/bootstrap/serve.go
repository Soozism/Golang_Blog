package bootstrap

import (
	"golang_blog/pkg/config"
	"golang_blog/pkg/html"
	"golang_blog/pkg/routing"
	"golang_blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
