package models

import "gopkg.in/mgo.v2/bson"

type DocumentType struct {
	Id   bson.ObjectId
	Name string
}
