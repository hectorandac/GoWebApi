package main

import (
	"github.com/go-martini/martini"
	"github.com/hectorandac/GoWebApi/handlers/korean_words"
	"github.com/hectorandac/GoWebApi/middlewares"
	"github.com/hectorandac/GoWebApi/models"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(middlewares.MongoDB())
	m.Use(render.Renderer())

	m.Get("/", korean_words.List)
	m.Get("/new", korean_words.AddEdit)
	m.Post("/korean_words", binding.Form(models.KoreanWord{}), korean_words.Add)
	m.Get("/korean_words/:_id", korean_words.Show)

	m.Run()
}
