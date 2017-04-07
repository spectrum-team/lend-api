package repos

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

var (
	assetRepo *Asset
)

func TestRepos(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Could not connect to the database. %s", err.Error())
	}

	db := session.DB("mini_biz")

	assetRepo = NewAssetRepo(db)

	//Tests
	t.Run("Insert Asset", insertAssetTest)
	t.Run("Find Assets", findAllAssetTest)
	t.Run("Find Asset By Id", findAssetByIdTest)
	t.Run("Update Asset", updateAssetTest)
}
