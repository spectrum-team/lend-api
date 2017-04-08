package controllers

import (
	"encoding/json"
	"lend-api/repos"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AssetController struct {
	repo *repos.Asset
}

func NewAssetController(db *mgo.Database) *AssetController {
	asset := new(AssetController)
	asset.repo = repos.NewAssetRepo(db)

	return asset
}

func (a *AssetController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		http.Error(w, "Invalid asset id", http.StatusBadRequest)
		return
	}

	req := bson.ObjectIdHex(id)

	asset, err := a.repo.FindById(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asset)
}
