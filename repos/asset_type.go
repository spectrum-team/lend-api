package repos

import (
	"lend-api/models"

	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AssetCategory struct {
	DB *mgo.Database
}

func NewAssetCategoryRepo(db *mgo.Database) *AssetCategory {
	assetCategory := new(AssetCategory)
	assetCategory.DB = db

	return assetCategory
}

// FindById looks for an asset by an id passed
func (a *AssetCategory) FindById(id bson.ObjectId) (*models.AssetCategory, error) {
	assetCategory := new(models.AssetCategory)

	err := a.DB.C("asset_category").FindId(id).One(assetCategory)
	if err != nil {
		return nil, err
	}

	return assetCategory, nil
}

func (a *AssetCategory) Find(query interface{}) ([]*models.AssetCategory, error) {
	assetCategorys := make([]*models.AssetCategory, 0)

	err := a.DB.C("asset_category").Find(query).All(&assetCategorys)
	if err != nil {
		return nil, err
	}

	return assetCategorys, nil
}

func (a *AssetCategory) Insert(assetCategory models.AssetCategory) error {
	if assetCategory.Name == "" {
		return errors.New("AssetCategory should have a name")
	}

	assetCategory.Id = bson.NewObjectId()

	return a.DB.C("asset_category").Insert(assetCategory)
}

func (a *AssetCategory) Update(id bson.ObjectId, assetCategory models.AssetCategory) error {
	if assetCategory.Name == "" {
		return errors.New("AssetCategory should have a name")
	}

	selector := bson.M{"_id": id}

	return a.DB.C("asset_category").Update(selector, assetCategory)
}
