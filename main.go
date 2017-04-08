package main

import (
	"lend-api/controllers"
	"net/http"
	"os"

	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getDBSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {

	db, err := getDBSession()
	if err != nil {
		log.Fatal(err)
	}

	asset := controllers.NewAssetController(db.DB("mini_biz"))

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/asset/{id}", asset.FindById).Methods("GET")

	listen := ":9000"

	if err := http.ListenAndServe(listen, handlers.CombinedLoggingHandler(os.Stdout, router)); err != nil {
		log.Fatal(err)
	}
}
