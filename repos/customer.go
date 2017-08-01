package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	DB *mgo.Database
}

func NewCustomerRepo(db *mgo.Database) *Customer {
	customer := new(Customer)
	customer.DB = db

	return customer
}

// FindById looks for an customer by an id passed
func (a *Customer) FindById(id bson.ObjectId) (*models.Customer, error) {
	customer := new(models.Customer)

	err := a.DB.C("customer").FindId(id).One(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (a *Customer) Find(query interface{}) ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)

	err := a.DB.C("customer").Find(query).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (a *Customer) Insert(customer models.Customer) error {
	customer.Id = bson.NewObjectId()
	return a.DB.C("customer").Insert(customer)
}

func (a *Customer) Update(id bson.ObjectId, customer models.Customer) error {

	selector := bson.M{"_id": id}

	return a.DB.C("customer").Update(selector, customer)
}
