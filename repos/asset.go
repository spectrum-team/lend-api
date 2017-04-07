package repos

import (
	"lend-api/models"

	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Asset struct {
	DB *mgo.Database
}

func NewAssetRepo(db *mgo.Database) *Asset {
	asset := new(Asset)
	asset.DB = db

	return asset
}

// FindById looks for an asset by an id passed
func (a *Asset) FindById(id bson.ObjectId) (*models.Asset, error) {
	asset := new(models.Asset)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("asset").FindId(id).One(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (a *Asset) Find(query interface{}) ([]*models.Asset, error) {
	assets := make([]*models.Asset, 0)

	err := a.DB.C("asset").Find(query).All(&assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (a *Asset) Insert(asset models.Asset) error {
	if asset.Name == "" {
		return errors.New("Asset should have a name")
	}

	if asset.UnitPrice < 1 {
		return errors.New("Asset should have a UnitPrice")
	}

	// if asset.Category == nil {
	// 	return errors.New("Asset should have a Category")
	// }

	return a.DB.C("asset").Insert(asset)
}

func (a *Asset) Update(id bson.ObjectId, asset models.Asset) error {
	if asset.Name == "" {
		return errors.New("Asset should have a name")
	}

	if asset.UnitPrice < 1 {
		return errors.New("Asset should have a UnitPrice")
	}

	// if asset.Category == nil {
	// 	return errors.New("Asset should have a Category")
	// }

	selector := bson.M{"_id": id}

	return a.DB.C("asset").Update(selector, asset)
}
