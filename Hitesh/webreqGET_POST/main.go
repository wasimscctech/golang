package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content)) //the above 4 lines are same as this print line
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	// requestBody := strings.NewReader(``) //we can create any new reader format inside newreader parantheses using `` sign
	requestBody := strings.NewReader(`
		{
			"MyName":"Wasim"
			"Age":"26"
			"email":"wasimshaikh@golang.dev"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody) //post needs three parameters. to get the content type, new request>>headers>>response headers>>content type
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Wasim")
	data.Add("lastname", "Shaikh")
	data.Add("email", "Wasimahmad@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
