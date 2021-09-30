package middlewares

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Templates() martini.Handler {
	return render.Renderer(render.Options{
		Layout: "layouts/default",
	})
}
