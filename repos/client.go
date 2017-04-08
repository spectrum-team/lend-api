package repos

import (
	"lend-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	DB *mgo.Database
}

func NewClientRepo(db *mgo.Database) *Client {
	client := new(Client)
	client.DB = db

	return client
}

// FindById looks for an client by an id passed
func (a *Client) FindById(id bson.ObjectId) (*models.Client, error) {
	client := new(models.Client)

	// query := bson.M{"_id": id.Hex()}

	err := a.DB.C("client").FindId(id).One(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (a *Client) Find(query interface{}) ([]*models.Client, error) {
	clients := make([]*models.Client, 0)

	err := a.DB.C("client").Find(query).All(&clients)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (a *Client) Insert(client models.Client) error {
	return a.DB.C("client").Insert(client)
}

func (a *Client) Update(id bson.ObjectId, client models.Client) error {

	selector := bson.M{"_id": id}

	return a.DB.C("client").Update(selector, client)
}
