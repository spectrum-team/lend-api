package models

import "gopkg.in/mgo.v2/bson"

type TransactionType struct {
	Id   bson.ObjectId
	Name string
}
