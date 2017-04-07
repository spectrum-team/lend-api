package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	Id             bson.ObjectId
	EntryDate      time.Time
	ExpirationDate time.Time
	Amount         float64
	Taxes          float64
	Client         *Client
	Details        DocumentDetails
}

type DocumentDetails struct {
	Asset    *Asset
	Quantity int64
}
