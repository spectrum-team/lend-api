package repos

import (
	"lend-api/models"
	"testing"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var (
	assetId bson.ObjectId
)

func insertAssetTest(t *testing.T) {
	obj := models.Asset{
		Id:         bson.NewObjectId(),
		Name:       "Honda Civic 2010",
		UnitPrice:  10.50,
		IncludeTax: false,
		Quantity:   10,
		// Category:   &models.AssetCategory{},
	}

	err := assetRepo.Insert(obj)
	if err != nil {
		t.Errorf("there was an error saving the asset. %s", err.Error())
	}

	obj.Id = bson.NewObjectId()
	obj.Name = "Toyota Corolla"

	err = assetRepo.Insert(obj)
	if err != nil {
		t.Errorf("there was an error saving the asset. %s", err.Error())
	}
}

func findAllAssetTest(t *testing.T) {
	assets, err := assetRepo.Find(nil)
	if err != nil {
		t.Errorf("There was an error looking for the asset. %s", err.Error())
	}

	fmt.Println(assets[0].Id)
	assetId = assets[0].Id
}

func findAssetByIdTest(t *testing.T) {
	fmt.Println(assetId)
	//lol it says ass xD
	ass, err := assetRepo.FindById(assetId)
	if err != nil {
		t.Errorf("There was an error looking for the asset. %s", err.Error())
	}

	//LOL it says ass again xd
	if ass != nil && ass.Id != assetId {
		t.Errorf("Found the wrong Asset")
	}
}

func updateAssetTest(t *testing.T) {
	obj := models.Asset{
		Id:         assetId,
		Name:       "Plancha a Vapor",
		UnitPrice:  1000.50,
		IncludeTax: false,
		Quantity:   10,
	}

	err := assetRepo.Update(assetId, obj)
	if err != nil {
		t.Errorf("There was an error updating the asset. %s", err.Error())
	}

	ass, err := assetRepo.FindById(assetId)
	if err != nil {
		t.Errorf("There was an error getting the asset. %s", err.Error())
	}

	if ass != nil && ass.Name != "Plancha a Vapor" {
		t.Errorf("The update did not happened")
	}
}
