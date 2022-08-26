package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://facebook.com"

func main() {
	fmt.Println("Web Request")
	response, err := http.Get(url) //simple http request. Internet connection must be there

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of the response %T\n", response) //the pointer type (the type is: *http.Response)

	defer response.Body.Close() //It is caller's responsibility to close the connection.

	databytes, err := ioutil.ReadAll(response.Body) //majority of the time to read the file we will use the "ioutil"

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
