package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	//running a server
	// http.ListenAndServe(":4000", r) this code may throw an error so we will handle the error using log.fatal. we can also use comma okay syntax. but mostly in case of web we will use log.fatla
	log.Fatal(http.ListenAndServe(":4000", r)) //4000 is a port and r is a router. if something fails it locks the value
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) { //a common syntax we will be using for designing a backend
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
