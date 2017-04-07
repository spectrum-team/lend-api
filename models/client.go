package models

import "gopkg.in/mgo.v2/bson"

type Client struct {
	Id        bson.ObjectId
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
