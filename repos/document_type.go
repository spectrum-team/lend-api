package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DocumentType struct {
	DB *mgo.Database
}

func NewDocumentTypeRepo(db *mgo.Database) *DocumentType {
	documentType := new(DocumentType)
	documentType.DB = db

	return documentType
}

// FindById looks for an documentType by an id passed
func (a *DocumentType) FindById(id bson.ObjectId) (*models.DocumentType, error) {
	documentType := new(models.DocumentType)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("document_type").FindId(id).One(documentType)
	if err != nil {
		return nil, err
	}

	return documentType, nil
}

func (a *DocumentType) Find(query interface{}) ([]*models.DocumentType, error) {
	documentTypes := make([]*models.DocumentType, 0)

	err := a.DB.C("document_type").Find(query).All(&documentTypes)
	if err != nil {
		return nil, err
	}

	return documentTypes, nil
}

func (a *DocumentType) Insert(documentType models.DocumentType) error {
	documentType.Id = bson.NewObjectId()
	return a.DB.C("document_type").Insert(documentType)
}

func (a *DocumentType) Update(id bson.ObjectId, documentType models.DocumentType) error {

	selector := bson.M{"_id": id}

	return a.DB.C("document_type").Update(selector, documentType)
}
