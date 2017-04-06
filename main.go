package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// func openDBConnection () error {

// }

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", healthCheck).Methods("GET")

	listen := "9000"

	http.ListenAndServe(listen, handlers.CombinedLoggingHandler(os.Stdout, router))
}
