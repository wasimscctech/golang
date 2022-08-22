package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerRouting() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", CreateEmployee).Methods("POST")
	// r.HandlerFunc("/employee", CreateEmployee()).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
