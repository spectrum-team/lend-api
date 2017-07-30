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

// DocumentController is the struct used to abstract all the functions in the Document API
type DocumentController struct {
	repo *repos.Document
}

// NewDocumentController function returns a pointer to a new DocumentController
func NewDocumentController(db *mgo.Database) *DocumentController {
	document := new(DocumentController)
	document.repo = repos.NewDocumentRepo(db)

	return document
}

// FindByID Gets an document by an ID passed on the url
func (a *DocumentController) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errResponse := &models.ErrorResponse{}

	id := vars["id"]

	if id == "" {
		errResponse.ApplicationMessage = ""
		errResponse.UserMessage = "No document ID specified"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	req := bson.ObjectIdHex(id)

	document, err := a.repo.FindById(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to get the document"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(document)
}

// Find receives a JSON containing the query to be made to the database
func (a *DocumentController) Find(w http.ResponseWriter, r *http.Request) {
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

	documments, err := a.repo.Find(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to find documments"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(documments)
}

// Insert creates a new Asset
func (a *DocumentController) Insert(w http.ResponseWriter, r *http.Request) {
	req := &models.Document{}
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
		errResponse.UserMessage = "Error trying to create the document"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}

// Update refreshes a document set using the ID passed in the URL and the json body
func (a *DocumentController) Update(w http.ResponseWriter, r *http.Request) {
	req := &models.Document{}
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
		errResponse.UserMessage = "Error trying to update the document"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}
