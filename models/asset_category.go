package models

import "gopkg.in/mgo.v2/bson"

type AssetCategory struct {
	Id   bson.ObjectId
	Name string
}
