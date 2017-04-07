package main

import (
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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	db, err := getDBSession()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", healthCheck).Methods("GET")

	listen := "9000"

	http.ListenAndServe(listen, handlers.CombinedLoggingHandler(os.Stdout, router))
}
