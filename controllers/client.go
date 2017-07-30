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

// ClientController is the struct used to abstract all the functions in the Client API
type ClientController struct {
	repo *repos.Client
}

// NewClientController function returns a pointer to a new ClientController
func NewClientController(db *mgo.Database) *ClientController {
	client := new(ClientController)
	client.repo = repos.NewClientRepo(db)

	return client
}

// FindByID Gets an Client by an ID passed on the url
func (a *ClientController) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errResponse := &models.ErrorResponse{}

	id := vars["id"]

	if id == "" {
		errResponse.ApplicationMessage = ""
		errResponse.UserMessage = "No Client ID specified"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	req := bson.ObjectIdHex(id)

	client, err := a.repo.FindById(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to get the Client"
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// Find receives a JSON containing the query to be made to the database
func (a *ClientController) Find(w http.ResponseWriter, r *http.Request) {
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

	clients, err := a.repo.Find(req)
	if err != nil {
		errResponse.ApplicationMessage = err.Error()
		errResponse.UserMessage = "Error trying to find clients"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(clients)
}

// Insert creates a new Asset
func (a *ClientController) Insert(w http.ResponseWriter, r *http.Request) {
	req := &models.Client{}
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
		errResponse.UserMessage = "Error trying to create the client"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}

// Update refreshes a client set using the ID passed in the URL and the json body
func (a *ClientController) Update(w http.ResponseWriter, r *http.Request) {
	req := &models.Client{}
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
		errResponse.UserMessage = "Error trying to update the client"
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ok")
}
