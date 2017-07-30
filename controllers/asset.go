package controllers

import (
	"encoding/json"
	"lend-api/models"
	"lend-api/repos"
	"net/http"

	"io/ioutil"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AssetController is the struct used to abstract all the functions in the Asset API
type AssetController struct {
	repo *repos.Asset
}

// NewAssetController function returns a pointer to a new AssetController
func NewAssetController(db *mgo.Database) *AssetController {
	asset := new(AssetController)
	asset.repo = repos.NewAssetRepo(db)

	return asset
}

// FindByID Gets an Asset by an ID passed on the url
func (a *AssetController) FindByID(w http.ResponseWriter, r *http.Request) {
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

// Find receives a JSON containing the query to be made to the database
func (a *AssetController) Find(w http.ResponseWriter, r *http.Request) {
	var req interface{}
	errResponse := &models.ErrorResponse{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	assets, err := a.repo.Find(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to find assets"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(assets)
}

// Insert creates a new Asset
func (a *AssetController) Insert(w http.ResponseWriter, r *http.Request) {
	req := &models.Asset{}
	errResponse := &models.ErrorResponse{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = a.repo.Insert(*req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to create the asset"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}

// Update refreshes an asset using the ID passed in the URL and the json body
func (a *AssetController) Update(w http.ResponseWriter, r *http.Request) {
	req := &models.Asset{}
	errResponse := &models.ErrorResponse{}
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		errResponse.ApplicationMessage = ""
		errResponse.UserMessage = "ID not specified"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	bsonID := bson.ObjectIdHex(id)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error parsing json"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	err = a.repo.Update(bsonID, *req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to update the asset"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}
