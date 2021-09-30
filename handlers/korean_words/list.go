package korean_words

import (
	"github.com/go-martini/martini"
	"github.com/hectorandac/GoWebApi/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

/**
 * List
 */

func List(r render.Render, params martini.Params, db *mgo.Database) {

	var korean_words []models.KoreanWord

	err := db.C("korean_words").Find(nil).All(&korean_words)

	if err != nil {
		r.Error(400)
	}

	r.HTML(200, "korean_words/list", korean_words)
}
