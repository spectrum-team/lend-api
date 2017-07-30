package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	DB *mgo.Database
}

func NewDocumentRepo(db *mgo.Database) *Document {
	document := new(Document)
	document.DB = db

	return document
}

// FindById looks for an document by an id passed
func (a *Document) FindById(id bson.ObjectId) (*models.Document, error) {
	document := new(models.Document)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("document").FindId(id).One(document)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (a *Document) Find(query interface{}) ([]*models.Document, error) {
	documents := make([]*models.Document, 0)

	err := a.DB.C("document").Find(query).All(&documents)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (a *Document) Insert(document models.Document) error {
	document.Id = bson.NewObjectId()
	return a.DB.C("document").Insert(document)
}

func (a *Document) Update(id bson.ObjectId, document models.Document) error {

	selector := bson.M{"_id": id}

	return a.DB.C("document").Update(selector, document)
}
