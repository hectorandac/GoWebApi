package models

import "gopkg.in/mgo.v2/bson"

type KoreanWord struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value       string        `json:"value" form:"value" binding:"required" bson:"value"`
	Description string        `json:"description" form:"description" binding:"required" bson:"description"`
	CreatedOn   int64         `json:"created_on" bson:"created_on"`
}
