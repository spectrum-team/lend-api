package models

import "gopkg.in/mgo.v2/bson"

type Asset struct {
	Id         bson.ObjectId `json:"id" bson:"_id"` //Need to do this json/bson thing so MongoDB does not create two Ids (id and _id)
	Name       string
	UnitPrice  float64
	IncludeTax bool
	Quantity   int64
	Category   *AssetCategory
}
