package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { //request is something user sends to the server and response is something server sends back to the user.
	if r.URL.Path != "/hello" { //if the path is not /hello then we will return an error
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//when type /hello then bydefault it is get method
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound) //if the request is not GET then it will return an error
		return
	}
	fmt.Fprintf(w, "Hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request is successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //we are telling go lang to check the static directory. and golang knows that it has to look for index.html file
	fmt.Printf("type %T", fileServer)
	http.Handle("/", fileServer) //we are handling the root route and sending to the file handler we have created above (fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler) // we will handle hello route which will call helloHandler function.

	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
