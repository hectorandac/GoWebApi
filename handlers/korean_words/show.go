package korean_words

import (
	"github.com/go-martini/martini"
	"github.com/hectorandac/GoWebApi/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/**
 * Show
 */

func Show(params martini.Params, r render.Render, db *mgo.Database) {

	korean_word := models.KoreanWord{}
	oId := bson.ObjectIdHex(params["_id"])

	err := db.C("korean_words").FindId(oId).One(&korean_word)

	if err != nil {
		r.HTML(400, "400", err)
	}

	r.HTML(200, "korean_words/show", korean_word)
}
