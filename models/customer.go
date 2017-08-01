package models

import "gopkg.in/mgo.v2/bson"

// Customer defines the structure that will be used for customers.
type Customer struct {
	Id        bson.ObjectId
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Addresses []*Address
}

// Address define the structure of the addresses we will have for customers
// FullAddress define everthing related to the address (street, number, apt, floor, etc)
// Neighborhood defines the neighborhood
// Moreindications is used to define more information that can be helpful in case of shippings
type Address struct {
	Description     string
	Neighborhood    string
	MoreIndications *string
}
