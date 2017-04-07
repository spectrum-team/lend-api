package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Id             bson.ObjectId
	Asset          *Asset
	Type           *TransactionType
	Price          float64
	EntryDate      time.Time
	ExpirationDate *time.Time
}
