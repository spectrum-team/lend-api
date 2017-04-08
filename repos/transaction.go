package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	DB *mgo.Database
}

func NewTransactionRepo(db *mgo.Database) *Transaction {
	transaction := new(Transaction)
	transaction.DB = db

	return transaction
}

// FindById looks for an transaction by an id passed
func (a *Transaction) FindById(id bson.ObjectId) (*models.Transaction, error) {
	transaction := new(models.Transaction)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("transaction").FindId(id).One(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (a *Transaction) Find(query interface{}) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)

	err := a.DB.C("transaction").Find(query).All(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (a *Transaction) Insert(transaction models.Transaction) error {
	return a.DB.C("transaction").Insert(transaction)
}

func (a *Transaction) Update(id bson.ObjectId, transaction models.Transaction) error {

	selector := bson.M{"_id": id}

	return a.DB.C("transaction").Update(selector, transaction)
}
