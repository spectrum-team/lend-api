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
	session, err := mgo.Dial("mongodb://tavomoya:tavomoya@ds143030.mlab.com:43030/mini_biz")
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
	assetCategory := controllers.NewAssetCategoryController(db.DB("mini_biz"))
	client := controllers.NewClientController(db.DB("mini_biz"))
	document := controllers.NewDocumentController(db.DB("mini_biz"))
	documentType := controllers.NewDocumentTypeController(db.DB("mini_biz"))
	transaction := controllers.NewTransactionController(db.DB("mini_biz"))
	transactionType := controllers.NewTransactionTypeController(db.DB("mini_biz"))
	customer := controllers.NewCustomerController(db.DB("mini_biz"))

	defer db.Close()

	router := mux.NewRouter()

	// Asset Routes
	router.HandleFunc("/asset/{id}", asset.FindByID).Methods("GET")
	router.HandleFunc("/asset/find", asset.Find).Methods("POST")
	router.HandleFunc("/asset", asset.Insert).Methods("POST")
	router.HandleFunc("/asset", asset.Update).Methods("PUT")

	// AssetCategory Routes
	router.HandleFunc("/assetcategory/{id}", assetCategory.FindByID).Methods("GET")
	router.HandleFunc("/assetcategory/find", assetCategory.Find).Methods("POST")
	router.HandleFunc("/assetcategory", assetCategory.Insert).Methods("POST")
	router.HandleFunc("/assetcategory", assetCategory.Update).Methods("PUT")

	// Client Routes
	router.HandleFunc("/client/{id}", client.FindByID).Methods("GET")
	router.HandleFunc("/client/find", client.Find).Methods("POST")
	router.HandleFunc("/client", client.Insert).Methods("POST")
	router.HandleFunc("/client", client.Update).Methods("PUT")

	// Document Routes
	router.HandleFunc("/document/{id}", document.FindByID).Methods("GET")
	router.HandleFunc("/document/find", document.Find).Methods("POST")
	router.HandleFunc("/document", document.Insert).Methods("POST")
	router.HandleFunc("/document", document.Update).Methods("PUT")

	// DocumentType Routes
	router.HandleFunc("/documenttype/{id}", documentType.FindByID).Methods("GET")
	router.HandleFunc("/documenttype/find", documentType.Find).Methods("POST")
	router.HandleFunc("/documenttype", documentType.Insert).Methods("POST")
	router.HandleFunc("/documenttype", documentType.Update).Methods("PUT")

	// Transaction Routes
	router.HandleFunc("/transaction/{id}", transaction.FindByID).Methods("GET")
	router.HandleFunc("/transaction/find", transaction.Find).Methods("POST")
	router.HandleFunc("/transaction", transaction.Insert).Methods("POST")
	router.HandleFunc("/transaction", transaction.Update).Methods("PUT")

	// TransactionType Routes
	router.HandleFunc("/transactiontype/{id}", transactionType.FindByID).Methods("GET")
	router.HandleFunc("/transactiontype/find", transactionType.Find).Methods("POST")
	router.HandleFunc("/transactiontype", transactionType.Insert).Methods("POST")
	router.HandleFunc("/transactiontype", transactionType.Update).Methods("PUT")

	// Customer Routes
	router.HandleFunc("/customer/{id}", customer.FindByID).Methods("GET")
	router.HandleFunc("/customer/find", customer.Find).Methods("POST")
	router.HandleFunc("/customer", customer.Insert).Methods("POST")
	router.HandleFunc("/customer", customer.Update).Methods("PUT")

	listen := os.Getenv("PORT")

	if listen == "" {
		listen = "9000"
	}

	if err := http.ListenAndServe(":"+listen, handlers.CombinedLoggingHandler(os.Stdout, router)); err != nil {
		log.Fatal(err)
	}
}
