package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TransactionType struct {
	DB *mgo.Database
}

func NewTransactionTypeRepo(db *mgo.Database) *TransactionType {
	transactionType := new(TransactionType)
	transactionType.DB = db

	return transactionType
}

// FindById looks for an transactionType by an id passed
func (a *TransactionType) FindById(id bson.ObjectId) (*models.TransactionType, error) {
	transactionType := new(models.TransactionType)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("transaction_type").FindId(id).One(transactionType)
	if err != nil {
		return nil, err
	}

	return transactionType, nil
}

func (a *TransactionType) Find(query interface{}) ([]*models.TransactionType, error) {
	transactionTypes := make([]*models.TransactionType, 0)

	err := a.DB.C("transaction_type").Find(query).All(&transactionTypes)
	if err != nil {
		return nil, err
	}

	return transactionTypes, nil
}

func (a *TransactionType) Insert(transactionType models.TransactionType) error {
	return a.DB.C("transaction_type").Insert(transactionType)
}

func (a *TransactionType) Update(id bson.ObjectId, transactionType models.TransactionType) error {

	selector := bson.M{"_id": id}

	return a.DB.C("transaction_type").Update(selector, transactionType)
}
