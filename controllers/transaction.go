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

// TransactionController is the struct used to abstract all the functions in the Transaction API
type TransactionController struct {
	repo *repos.Transaction
}

// NewTransactionController function returns a pointer to a new TransactionController
func NewTransactionController(db *mgo.Database) *TransactionController {
	transaction := new(TransactionController)
	transaction.repo = repos.NewTransactionRepo(db)

	return transaction
}

// FindByID Gets an transaction by an ID passed on the url
func (a *TransactionController) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errResponse := &models.ErrorResponse{}

	id := vars["id"]

	if id == "" {
		errResponse.ApplicationMessage = ""
		errResponse.UserMessage = "No transaction ID specified"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	req := bson.ObjectIdHex(id)

	transaction, err := a.repo.FindById(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to get the transaction"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// Find receives a JSON containing the query to be made to the database
func (a *TransactionController) Find(w http.ResponseWriter, r *http.Request) {
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

	transactions, err := a.repo.Find(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to find transactions"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(transactions)
}

// Insert creates a new Transaction
func (a *TransactionController) Insert(w http.ResponseWriter, r *http.Request) {
	req := &models.Transaction{}
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
		errResponse.UserMessage = "Error trying to create the transaction"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}

// Update refreshes a transaction set using the ID passed in the URL and the json body
func (a *TransactionController) Update(w http.ResponseWriter, r *http.Request) {
	req := &models.Transaction{}
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
		errResponse.UserMessage = "Error trying to update the transaction"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}
