package korean_words

import (
	"fmt"

	"github.com/hectorandac/GoWebApi/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

/**
 * Get add/edit form
 */

func AddEdit(r render.Render) {
	r.HTML(200, "korean_words/form", nil)
}

/**
 * Add
 */

func Add(korean_word models.KoreanWord, r render.Render, db *mgo.Database) {
	fmt.Println(korean_word)
	err := db.C("korean_words").Insert(korean_word)

	if err != nil {
		r.HTML(400, "400", err)
	} else {
		r.Redirect("/")
	}

}
