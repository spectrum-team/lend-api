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

// TransactionTypeController is the struct used to abstract all the functions in the TransactionType API
type TransactionTypeController struct {
	repo *repos.TransactionType
}

// NewTransactionTypeController function returns a pointer to a new TransactionTypeController
func NewTransactionTypeController(db *mgo.Database) *TransactionTypeController {
	transactionType := new(TransactionTypeController)
	transactionType.repo = repos.NewTransactionTypeRepo(db)

	return transactionType
}

// FindByID Gets an transactionType by an ID passed on the url
func (a *TransactionTypeController) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errResponse := &models.ErrorResponse{}

	id := vars["id"]

	if id == "" {
		errResponse.ApplicationMessage = ""
		errResponse.UserMessage = "No transactionType ID specified"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	req := bson.ObjectIdHex(id)

	transactionType, err := a.repo.FindById(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to get the transactionType"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactionType)
}

// Find receives a JSON containing the query to be made to the database
func (a *TransactionTypeController) Find(w http.ResponseWriter, r *http.Request) {
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

	transactionTypes, err := a.repo.Find(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to find transactionTypes"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(transactionTypes)
}

// Insert creates a new TransactionType
func (a *TransactionTypeController) Insert(w http.ResponseWriter, r *http.Request) {
	req := &models.TransactionType{}
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
		errResponse.UserMessage = "Error trying to create the transactionType"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}

// Update refreshes a transactionType set using the ID passed in the URL and the json body
func (a *TransactionTypeController) Update(w http.ResponseWriter, r *http.Request) {
	req := &models.TransactionType{}
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
		errResponse.UserMessage = "Error trying to update the transactionType"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}
